package storage

import (
	"context"
	"time"

	"github.com/STRockefeller/langdida-server/models/protomodels"
)

type Storage interface {
	ListCards(ctx context.Context, cardIndex []protomodels.CardIndex) ([]protomodels.Card, error)

	// arguments details:
	//  - needReview: true => need to review, false => all
	ListCardsWithConditions(ctx context.Context, needReview bool, language protomodels.Language) ([]protomodels.Card, error)

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

type ListCardsConditions struct {
	NeedReview bool // true: need to review, false: all
	Language   protomodels.Language
	Tag        string
}
