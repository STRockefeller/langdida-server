package instance

import (
	"context"
	"fmt"

	"github.com/STRockefeller/dictionaries"
	"github.com/STRockefeller/go-linq"
	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/storage"
)

type CardService struct {
	storage storage.Storage
}

func NewCardService(storage storage.Storage) *CardService {
	return &CardService{storage: storage}
}

func (service CardService) GetCard(ctx context.Context, condition protomodels.CardIndex) (protomodels.Card, error) {
	rep, err := service.storage.ListCards(ctx, []protomodels.CardIndex{condition})
	if err != nil {
		return protomodels.Card{}, err
	}
	if len(rep) == 1 {
		return rep[1], nil
	}
	return protomodels.Card{}, fmt.Errorf("more than 1 cards found")
}

func (service CardService) CreateCard(ctx context.Context, card protomodels.Card) error {
	return service.storage.CreateCard(ctx, card)
}

func (service CardService) EditCard(ctx context.Context, card protomodels.Card) error {
	return service.storage.UpdateCard(ctx, card)
}

func (service CardService) ListCardsShouldBeReviewed(ctx context.Context, language protomodels.Language) ([]protomodels.Card, error) {
	return service.storage.ListCardsWithConditions(ctx, true, language)
}

func (service CardService) ListCardsByLabelsAndLanguage(ctx context.Context, labels []string, language protomodels.Language) ([]protomodels.Card, error) {
	cards, err := service.storage.ListCardsWithConditions(ctx, false, language)
	if err != nil {
		return nil, err
	}
	return linq.NewLinq(cards).Where(func(card protomodels.Card) bool {
		for _, label := range labels {
			if !linq.NewLinq(card.Labels).Contains(label) {
				return false
			}
		}
		return true
	}).ToSlice(), nil
}

func (service CardService) SearchWithDictionary(ctx context.Context, cardIndex protomodels.CardIndex) ([]string, error) {
	switch cardIndex.GetLanguage() {
	case protomodels.Language_ENGLISH:
		result, err := dictionaries.NewEnglishDictionary().Search(cardIndex.GetName())
		if err != nil {
			return nil, err
		}
		return result.ListAllMeanings(), nil

	case protomodels.Language_JAPANESE:
		result, err := dictionaries.NewJapaneseDictionary().Search(cardIndex.GetName())
		return result.ListAllMeanings(), err

	default:
		return nil, fmt.Errorf("unsupported language")
	}
}
