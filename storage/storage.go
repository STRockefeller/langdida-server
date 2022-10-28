package storage

import (
	"context"
	"time"

	"github.com/STRockefeller/langdida-server/protomodels"
)

type Storage interface {
	ListCards(ctx context.Context, cardIndex []protomodels.CardIndex) ([]protomodels.Card, error)
	// upsert to logs NewCards++
	CreateCard(ctx context.Context, card protomodels.Card) error
	// zero values will NOT been updated
	UpdateCard(ctx context.Context, card protomodels.Card) error
	DeleteCard(ctx context.Context, cardIndex protomodels.CardIndex) error

	GetLog(ctx context.Context, date time.Time) (protomodels.Log, error)
	ListLogs(ctx context.Context, from time.Time, until time.Time) ([]protomodels.Log, error)

	// upsert to logs ReviewedCards++
	// update card review date
	ReviewCard(ctx context.Context, cardIndex protomodels.CardIndex) error
}
