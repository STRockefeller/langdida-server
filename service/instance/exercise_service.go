package instance

import (
	"context"

	"github.com/STRockefeller/go-linq"
	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/storage"
	problemGenerator "github.com/STRockefeller/problems"
)

type ExerciseService struct {
	storage storage.Storage
}

func NewExerciseService(storage storage.Storage) *ExerciseService {
	return &ExerciseService{storage: storage}
}

func (e ExerciseService) CreateChoiceProblems(ctx context.Context, cards []protomodels.CardIndex) (problems []string, answers []string, err error) {
	fullCards, err := e.storage.ListCards(ctx, cards)
	if err != nil {
		return nil, nil, err
	}
	muiltiChoiceProblems := problemGenerator.GenerateMultiChoiceProblems(linq.Select(linq.NewLinq(fullCards), func(card protomodels.Card) problemGenerator.FlashCard {
		return problemGenerator.FlashCard{
			Word:        card.Index.Name,
			Sentences:   card.ExampleSentences,
			Definitions: card.Explanations,
		}
	}))

	for _, problem := range muiltiChoiceProblems {
		question := problem.Question + "\n"
		for i, choice := range problem.Choices {
			question += string(rune(i+'A')) + ". " + choice + "\n"
		}
		problems = append(problems, question)
		answers = append(answers, problem.Answer)
	}
	return
}

func (e ExerciseService) CreateFillingProblems(ctx context.Context, cards []protomodels.CardIndex) (problems []string, answers []string, err error) {
	fullCards, err := e.storage.ListCards(ctx, cards)
	if err != nil {
		return nil, nil, err
	}
	fprob := problemGenerator.GenerateFillInTheBlankProblems(linq.Select(linq.NewLinq(fullCards), func(card protomodels.Card) problemGenerator.FlashCard {
		return problemGenerator.FlashCard{
			Word:        card.Index.Name,
			Sentences:   card.ExampleSentences,
			Definitions: card.Explanations,
		}
	}))
	for _, problem := range fprob {
		problems = append(problems, problem.Question)
		answers = append(answers, problem.Answer)
	}
	return
}
