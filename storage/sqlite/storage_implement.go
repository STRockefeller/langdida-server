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
	"gorm.io/gorm/clause"

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
	result, err := storage.cardTable().
		WhereRaw(glinq.NewQueryString(`(name, language) IN ?`,
			linq.
				Select(cardIndex, func(c protomodels.CardIndex) [2]any { return [2]any{c.Name, c.Language} }).
				ToSlice()),
		).Find(ctx)
	if err != nil {
		return nil, err
	}
	return convertCardModels(result), nil
}

func (storage Storage) ListCardIndices(ctx context.Context) ([]protomodels.CardIndex, error) {
	card, err := storage.cardTable().SelectRaw([]string{"name", "language"}).Find(ctx)
	if err != nil {
		return nil, err
	}
	return linq.Select(card, func(c gm.Card) protomodels.CardIndex {
		return protomodels.CardIndex{
			Name:     c.Name,
			Language: c.Language,
		}
	}), nil
}

// arguments details:
//  - needReview: true => need to review, false => all
func (storage Storage) ListCardsWithConditions(ctx context.Context, conditions storage.ListCardsConditions) ([]protomodels.Card, error) {
	result, err := listCardsFilters(conditions, storage.cardTable()).Find(ctx)
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
	return db.WhereRaw(glinq.NewQueryString(`language = ?`, language))
}

func filterReviewDate(date itime.UnixTime, db glinq.DB[gm.Card]) glinq.DB[gm.Card] {
	return db.WhereRaw(glinq.NewQueryString(`review_date < ?`, date))
}

func filterLabel(label string, db glinq.DB[gm.Card]) glinq.DB[gm.Card] {
	return db.WhereRaw(glinq.NewQueryString(`labels LIKE ?`, "%"+label+"%"))
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
		if _, err := storage.logTable().
			Where(gm.Log{Date: log.Date}).
			Updates(ctx, log.WithNewCard()); err != nil {
			return err
		}
	}

	return storage.cardTable().Create(ctx, gm.NewCard(card))
}

// zero values will NOT been updated
func (storage Storage) UpdateCard(ctx context.Context, card protomodels.Card) error {
	_, err := storage.cardTable().WhereRaw(glinq.NewQueryString(`name = ? AND language = ?`, card.Index.Name, card.Index.Language)).Updates(ctx, gm.NewCard(card))
	return err
}

func (storage Storage) DeleteCard(ctx context.Context, cardIndex protomodels.CardIndex) error {
	condition := gm.Card{
		Name:     cardIndex.Name,
		Language: cardIndex.Language,
	}
	_, err := storage.cardTable().Where(condition).Delete(ctx)
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
	result, err := storage.logTable().
		WhereRaw(glinq.NewQueryString(`date >= ? AND date <= ?`, dayFormat(from), dayFormat(until))).
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
		if _, err := storage.logTable().
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

	_, err = storage.cardTable().
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

func (storage Storage) GetAssociations(ctx context.Context, cardIndex protomodels.CardIndex) (protomodels.RelatedCards, error) {
	rep, err := storage.relatedCardsTable().Where(gm.RelatedCards{
		Name:     cardIndex.Name,
		Language: cardIndex.Language,
	}).Take(ctx)
	if err != nil {
		return protomodels.RelatedCards{}, err
	}

	return rep.ToProtoModel(), nil
}

func (storage Storage) CreateAssociation(ctx context.Context, conditions storage.CreateAssociationConditions) error {
	cardIndex := protomodels.CardIndex{
		Name:     conditions.CardIndex.Name,
		Language: conditions.CardIndex.Language,
	}
	card := protomodels.RelatedCards{
		Index: &cardIndex,
	}

	relatedCardIndex := protomodels.CardIndex{
		Name:     conditions.RelatedCardIndex.Name,
		Language: conditions.RelatedCardIndex.Language,
	}
	relatedCard := protomodels.RelatedCards{
		Index: &relatedCardIndex,
	}

	switch conditions.Association {
	case protomodels.AssociationTypes_SYNONYMS:
		card.Synonyms = append(card.Synonyms, &relatedCardIndex)
		relatedCard.Synonyms = append(relatedCard.Synonyms, &cardIndex)
	case protomodels.AssociationTypes_ANTONYMS:
		card.Antonyms = append(card.Antonyms, &relatedCardIndex)
		relatedCard.Antonyms = append(relatedCard.Antonyms, &cardIndex)
	case protomodels.AssociationTypes_ORIGIN:
		card.Origin = &relatedCardIndex
		relatedCard.Derivatives = append(relatedCard.Derivatives, &cardIndex)
	case protomodels.AssociationTypes_DERIVATIVES:
		card.Derivatives = append(card.Derivatives, &relatedCardIndex)
		relatedCard.Origin = &cardIndex
	case protomodels.AssociationTypes_IN_OTHER_LANGUAGES:
		card.InOtherLanguages = append(card.InOtherLanguages, &relatedCardIndex)
		relatedCard.InOtherLanguages = append(relatedCard.InOtherLanguages, &cardIndex)
	case protomodels.AssociationTypes_OTHERS:
		card.Others = append(card.Others, &relatedCardIndex)
		relatedCard.Others = append(relatedCard.Others, &cardIndex)
	}

	if err := storage.relatedCardsTable().Upsert(ctx, clause.OnConflict{
		UpdateAll: true,
	}, turnEmptySlicesToNilInRelatedCards(gm.NewRelatedCards(card))); err != nil {
		return err
	}

	return storage.relatedCardsTable().Upsert(ctx, clause.OnConflict{
		UpdateAll: true,
	}, turnEmptySlicesToNilInRelatedCards(gm.NewRelatedCards(relatedCard)))
}

func turnEmptySlicesToNilInRelatedCards(rc gm.RelatedCards) gm.RelatedCards {
	return gm.RelatedCards{
		Name:             rc.Name,
		Language:         rc.Language,
		Synonyms:         turnEmptySlicesToNil(rc.Synonyms),
		Antonyms:         turnEmptySlicesToNil(rc.Antonyms),
		Origin:           rc.Origin,
		Derivatives:      turnEmptySlicesToNil(rc.Derivatives),
		InOtherLanguages: turnEmptySlicesToNil(rc.InOtherLanguages),
		Others:           turnEmptySlicesToNil(rc.Others),
	}
}

func turnEmptySlicesToNil(as gm.ArrayOfStrings) gm.ArrayOfStrings {
	if len(as) == 0 {
		return nil
	}
	return as
}

func (storage Storage) createLog(ctx context.Context, log gm.Log) error {
	return storage.logTable().Create(ctx, log)
}

func (storage Storage) getLog(ctx context.Context, date time.Time) (gm.Log, error) {
	return storage.logTable().Where(gm.Log{Date: dayFormat(date)}).Take(ctx)
}

func (storage Storage) streak(ctx context.Context, date time.Time) int32 {
	log, err := storage.logTable().
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
