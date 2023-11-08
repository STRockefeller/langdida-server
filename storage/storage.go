package storage

import (
	"context"
	"time"

	"github.com/STRockefeller/langdida-server/models/protomodels"
)

type Storage interface {
	// todo : refactor with builder pattern
	ListCards(ctx context.Context, req ListCardsRequest) ([]protomodels.Card, error)

	ListCardIndices(ctx context.Context) ([]protomodels.CardIndex, error)

	GetAssociations(ctx context.Context, cardIndex protomodels.CardIndex) (protomodels.RelatedCards, error)

	CreateAssociation(ctx context.Context, conditions CreateAssociationConditions) error

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

type CreateAssociationConditions struct {
	CardIndex        protomodels.CardIndex
	RelatedCardIndex protomodels.CardIndex
	Association      protomodels.AssociationTypes
}
