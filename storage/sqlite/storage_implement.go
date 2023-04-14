package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/STRockefeller/go-linq"
	glinq "github.com/STRockefeller/gorm-linq"
	itime "github.com/STRockefeller/langdida-server/internal/time"
	"github.com/STRockefeller/langdida-server/models/gormmodels"
	"github.com/STRockefeller/langdida-server/models/protomodels"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewStorage(dbPath string, migrateTables bool) *Storage {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to open sqlite db, path: %s", dbPath))
	}

	if migrateTables {
		if err := db.AutoMigrate(&gormmodels.Card{}); err != nil {
			panic(err)
		}
	}
	return &Storage{
		db: db,
	}
}

type Storage struct {
	db *gorm.DB
}

func convertCardModels(cards linq.Linq[gormmodels.Card]) []protomodels.Card {
	return linq.Select(cards, func(c gormmodels.Card) protomodels.Card {
		return c.ToProtoModel()
	})
}

func (storage Storage) ListCards(ctx context.Context, cardIndex []protomodels.CardIndex) ([]protomodels.Card, error) {
	result, err := glinq.NewDB[gormmodels.Card](storage.db).
		WhereRaw(`(name, language) IN ?`,
			linq.Select(
				linq.NewLinq(cardIndex),
				func(c protomodels.CardIndex) [2]any { return [2]any{c.Name, c.Language} },
			).ToSlice()).Find(ctx)
	if err != nil {
		return nil, err
	}
	return convertCardModels(result), nil
}

// arguments details:
//  - needReview: true => need to review, false => all
func (storage Storage) ListCardsWithConditions(ctx context.Context, needReview bool, language protomodels.Language) ([]protomodels.Card, error) {
	date := itime.NewFromTime(time.Now())
	result, err := glinq.NewDB[gormmodels.Card](storage.db).Where(gormmodels.Card{
		Language: language,
	}).WhereRaw(`review_date < ?`, date).Find(ctx)

	if err != nil {
		return nil, err
	}

	return convertCardModels(result), nil
}

// upsert to logs NewCards++
func (storage Storage) CreateCard(ctx context.Context, card protomodels.Card) error {
	return glinq.NewDB[gormmodels.Card](storage.db).Create(ctx, []gormmodels.Card{gormmodels.NewCard(card)})
}

// zero values will NOT been updated
func (storage Storage) UpdateCard(ctx context.Context, card protomodels.Card) error {
	_, err := glinq.NewDB[gormmodels.Card](storage.db).Updates(ctx, gormmodels.NewCard(card))
	return err
}

func (storage Storage) DeleteCard(ctx context.Context, cardIndex protomodels.CardIndex) error {
	cond := gormmodels.Card{
		Name:     cardIndex.Name,
		Language: cardIndex.Language,
	}
	_, err := glinq.NewDB[gormmodels.Card](storage.db).Where(cond).Delete(ctx, cond)
	return err
}

func (storage Storage) GetLog(ctx context.Context, date time.Time) (protomodels.Log, error) {
	panic("not implemented") // TODO: Implement
}

func (storage Storage) ListLogs(ctx context.Context, from time.Time, until time.Time) ([]protomodels.Log, error) {
	panic("not implemented") // TODO: Implement
}

// upsert to logs ReviewedCards++
// update card review date
func (storage Storage) ReviewCard(ctx context.Context, cardIndex protomodels.CardIndex) error {
	panic("not implemented") // TODO: Implement
}
