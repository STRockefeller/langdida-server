package sqlite

import (
	"context"
	"time"

	linq "github.com/STRockefeller/gorm-linq"
	"github.com/STRockefeller/langdida-server/models/protomodels"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewStorage(dbPath string) *Storage {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to open sqlite db")
	}
	return &Storage{
		db: db,
	}
}

type Storage struct {
	db *gorm.DB
}

func (storage Storage) ListCards(ctx context.Context, cardIndex []protomodels.CardIndex) ([]protomodels.Card, error) {
	panic("not implemented") // TODO: Implement
}

// arguments details:
//  - needReview: true => need to review, false => all
func (storage Storage) ListCardsWithConditions(ctx context.Context, needReview bool, language protomodels.Language) ([]protomodels.Card, error) {
	panic("not implemented") // TODO: Implement
}

// upsert to logs NewCards++
func (storage Storage) CreateCard(ctx context.Context, card protomodels.Card) error {
	return linq.NewDB[protomodels.Card](storage.db).Create(ctx, []protomodels.Card{card})
}

// zero values will NOT been updated
func (storage Storage) UpdateCard(ctx context.Context, card protomodels.Card) error {
	panic("not implemented") // TODO: Implement
}

func (storage Storage) DeleteCard(ctx context.Context, cardIndex protomodels.CardIndex) error {
	panic("not implemented") // TODO: Implement
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
