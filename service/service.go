package service

import (
	"context"

	"github.com/STRockefeller/langdida-server/models/protomodels"
)

type CardService interface {
	GetCard(ctx context.Context, condition protomodels.CardIndex) (protomodels.Card, error)
	CreateCard(ctx context.Context, card protomodels.Card) error
	// renew the review date
	EditCard(ctx context.Context, card protomodels.Card) error
	ListCardsShouldBeReviewed(ctx context.Context) ([]protomodels.Card, error)
	ListCardsByLabelsAndLanguage(ctx context.Context, labels []string, language protomodels.Language) ([]protomodels.Card, error)

	// return url
	SearchWithDictionary(ctx context.Context, cardIndex protomodels.CardIndex) (string, error)
}

type LogService interface {
	GetLogStatus(ctx context.Context) (LogStatus, error)
}

type ExerciseService interface {
	CreateChoiceProblems(ctx context.Context, cards protomodels.CardIndex) (problems []string, answers []string, err error)
	CreateFillingProblems(ctx context.Context, cards protomodels.CardIndex) (problems []string, answers []string, err error)
}

type IOService interface{}
