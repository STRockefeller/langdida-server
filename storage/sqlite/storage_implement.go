package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/STRockefeller/go-linq"
	glinq "github.com/STRockefeller/gorm-linq"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/STRockefeller/langdida-server/internal/review"
	itime "github.com/STRockefeller/langdida-server/internal/time"
	gm "github.com/STRockefeller/langdida-server/models/gormmodels"
	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/storage"
)

func NewStorage(dbPath string, migrateTables bool) *Storage {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to open sqlite db, path: %s", dbPath))
	}
	zap.L().Info("successfully opened sqlite db", zap.String("path", dbPath))

	if migrateTables {
		if err := db.AutoMigrate(&gm.Card{}, &gm.Log{}); err != nil {
			panic(err)
		}
		zap.L().Info(`auto-migrate tables completed`)
	}
	return &Storage{
		db: db,
	}
}

type Storage struct {
	db *gorm.DB
}

func convertCardModels(cards linq.Linq[gm.Card]) []protomodels.Card {
	return linq.Select(cards, func(c gm.Card) protomodels.Card {
		return c.ToProtoModel()
	})
}

func (storage Storage) ListCards(ctx context.Context, cardIndex []protomodels.CardIndex) ([]protomodels.Card, error) {
	result, err := glinq.NewDB[gm.Card](storage.db).
		WhereRaw(`(name, language) IN ?`,
			linq.
				Select(cardIndex, func(c protomodels.CardIndex) [2]any { return [2]any{c.Name, c.Language} }).
				ToSlice()).Find(ctx)
	if err != nil {
		return nil, err
	}
	return convertCardModels(result), nil
}

// arguments details:
//  - needReview: true => need to review, false => all
func (storage Storage) ListCardsWithConditions(ctx context.Context, conditions storage.ListCardsConditions) ([]protomodels.Card, error) {
	result, err := listCardsFilters(conditions, glinq.NewDB[gm.Card](storage.db)).Find(ctx)
	if err != nil {
		return nil, err
	}

	return convertCardModels(result), nil
}

func listCardsFilters(conditions storage.ListCardsConditions, db glinq.DB[gm.Card]) glinq.DB[gm.Card] {
	if conditions.NeedReview {
		date := itime.NewFromTime(time.Now())
		db = filterReviewDate(date, db)
	}
	if conditions.Label != "" {
		db = filterLabel(conditions.Label, db)
	}
	if conditions.Language != nil {
		db = filterLanguage(*conditions.Language, db)
	}
	return db
}

func filterLanguage(language protomodels.Language, db glinq.DB[gm.Card]) glinq.DB[gm.Card] {
	return db.WhereRaw(`language = ?`, language)
}

func filterReviewDate(date itime.UnixTime, db glinq.DB[gm.Card]) glinq.DB[gm.Card] {
	return db.WhereRaw(`review_date < ?`, date)
}

func filterLabel(label string, db glinq.DB[gm.Card]) glinq.DB[gm.Card] {
	return db.WhereRaw(`labels LIKE ?`, "%"+label+"%")
}

// upsert to logs NewCards++
func (storage Storage) CreateCard(ctx context.Context, card protomodels.Card) error {
	now := time.Now()
	log, err := storage.getLog(ctx, now)
	if err != nil {
		log = gm.NewDefaultLog(storage.streak(ctx, now) + 1).WithNewCard()
		if err := storage.createLog(ctx, log); err != nil {
			return err
		}
	} else {
		if _, err := glinq.NewDB[gm.Log](storage.db).
			Where(gm.Log{Date: log.Date}).
			Updates(ctx, log.WithNewCard()); err != nil {
			return err
		}
	}

	return glinq.NewDB[gm.Card](storage.db).Create(ctx, gm.NewCard(card))
}

// zero values will NOT been updated
func (storage Storage) UpdateCard(ctx context.Context, card protomodels.Card) error {
	_, err := glinq.NewDB[gm.Card](storage.db).WhereRaw(`name = ? AND language = ?`, card.Index.Name, card.Index.Language).Updates(ctx, gm.NewCard(card))
	return err
}

func (storage Storage) DeleteCard(ctx context.Context, cardIndex protomodels.CardIndex) error {
	condition := gm.Card{
		Name:     cardIndex.Name,
		Language: cardIndex.Language,
	}
	_, err := glinq.NewDB[gm.Card](storage.db).Where(condition).Delete(ctx, condition)
	return err
}

func (storage Storage) GetLog(ctx context.Context, date time.Time) (protomodels.Log, error) {
	log, err := storage.getLog(ctx, date)
	if err != nil {
		return protomodels.Log{}, err
	}
	return log.ToProtoModel(), nil
}

func (storage Storage) ListLogs(ctx context.Context, from time.Time, until time.Time) ([]protomodels.Log, error) {
	result, err := glinq.NewDB[gm.Log](storage.db).
		WhereRaw(`date >= ? AND date <= ?`, dayFormat(from), dayFormat(until)).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	return linq.Select(result, func(l gm.Log) protomodels.Log { return l.ToProtoModel() }).ToSlice(), nil
}

// upsert to logs ReviewedCards++
// update card review date
func (storage Storage) ReviewCard(ctx context.Context, cardIndex protomodels.CardIndex) error {
	now := time.Now()
	log, err := storage.getLog(ctx, now)
	if err != nil {
		log = gm.NewDefaultLog(storage.streak(ctx, now) + 1).WithReviewedCard()
		if err := storage.createLog(ctx, log); err != nil {
			return err
		}
	} else {
		if _, err := glinq.NewDB[gm.Log](storage.db).
			Where(gm.Log{Date: log.Date}).
			Updates(ctx, log.WithReviewedCard()); err != nil {
			return err
		}
	}

	// query familiarity
	var familiarity int32
	if err := storage.db.WithContext(ctx).Select(`familiarity`).Where(gm.Card{
		Name:     cardIndex.Name,
		Language: cardIndex.Language,
	}).Find(&familiarity).Error; err != nil {
		return err
	}

	_, err = glinq.NewDB[gm.Card](storage.db).
		Where(gm.Card{
			Name:     cardIndex.Name,
			Language: cardIndex.Language,
		}).
		Updates(ctx, gm.Card{
			ReviewDate:  itime.NewFromTime(review.NextReviewDate(familiarity)),
			Familiarity: familiarity + 1,
		})
	return err
}

func (storage Storage) createLog(ctx context.Context, log gm.Log) error {
	return glinq.NewDB[gm.Log](storage.db).Create(ctx, log)
}

func (storage Storage) getLog(ctx context.Context, date time.Time) (gm.Log, error) {
	return glinq.NewDB[gm.Log](storage.db).Where(gm.Log{Date: dayFormat(date)}).Take(ctx)
}

func (storage Storage) streak(ctx context.Context, date time.Time) int32 {
	log, err := glinq.NewDB[gm.Log](storage.db).
		Where(gm.Log{Date: dayFormat(oneDayBefore(date))}).
		Take(ctx)
	if err != nil {
		return 0
	}
	return log.Streak
}

func oneDayBefore(t time.Time) time.Time {
	return t.Add(-24 * time.Hour)
}

func dayFormat(t time.Time) string {
	return t.Format("2006-01-02")
}
