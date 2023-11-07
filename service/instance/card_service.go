package instance

import (
	"context"
	"fmt"

	"github.com/STRockefeller/dictionaries"

	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/service"
	"github.com/STRockefeller/langdida-server/storage"
)

type CardService struct {
	storage storage.Storage
}

func (service CardService) ListCards(ctx context.Context, req storage.ListCardsRequest) ([]protomodels.Card, error) {
	return service.storage.ListCards(ctx, req)
}

func NewCardService(storage storage.Storage) service.CardService {
	return &CardService{storage: storage}
}

func (service CardService) GetCard(ctx context.Context, condition protomodels.CardIndex) (protomodels.Card, error) {
	rep, err := service.storage.ListCards(ctx, storage.NewListCardRequest().WhereCardIndexIn([]protomodels.CardIndex{condition}))
	if err != nil {
		return protomodels.Card{}, err
	}
	if len(rep) == 1 {
		return rep[0], nil
	}
	return protomodels.Card{}, fmt.Errorf("more than 1 cards found")
}

func (service CardService) CreateCard(ctx context.Context, card protomodels.Card) error {
	return service.storage.CreateCard(ctx, card)
}

func (service CardService) EditCard(ctx context.Context, card protomodels.Card) error {
	return service.storage.UpdateCard(ctx, card)
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

func (service CardService) ListIndices(ctx context.Context) ([]protomodels.CardIndex, error) {
	return service.storage.ListCardIndices(ctx)
}

func (service CardService) GetAssociations(ctx context.Context, cardIndex protomodels.CardIndex) (protomodels.RelatedCards, error) {
	return service.storage.GetAssociations(ctx, cardIndex)
}

func (service CardService) CreateAssociations(ctx context.Context, conditions storage.CreateAssociationConditions) error {
	return service.storage.CreateAssociation(ctx, conditions)
}
