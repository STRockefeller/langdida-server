package service

import (
	"context"

	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/storage"
)

type CardService interface {
	GetCard(ctx context.Context, condition protomodels.CardIndex) (protomodels.Card, error)
	CreateCard(ctx context.Context, card protomodels.Card) error
	// renew the review date
	EditCard(ctx context.Context, card protomodels.Card) error
	ListCards(ctx context.Context, req storage.ListCardsRequest) ([]protomodels.Card, error)
	ListIndices(ctx context.Context) ([]protomodels.CardIndex, error)
	GetAssociations(ctx context.Context, cardIndex protomodels.CardIndex) (protomodels.RelatedCards, error)
	CreateAssociations(ctx context.Context, conditions storage.CreateAssociationConditions) error

	// return url
	SearchWithDictionary(ctx context.Context, cardIndex protomodels.CardIndex) ([]string, error)
}

type LogService interface {
	GetLogStatus(ctx context.Context) (LogStatus, error)
}

type ExerciseService interface {
	CreateChoiceProblems(ctx context.Context, cards []protomodels.CardIndex) (problems []string, answers []string, err error)
	CreateFillingProblems(ctx context.Context, cards []protomodels.CardIndex) (problems []string, answers []string, err error)
}

type IOService interface {
	ImportFromURL(ctx context.Context, url string) (string, error)
}
