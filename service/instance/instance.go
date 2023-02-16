package instance

import (
	"context"
	"fmt"

	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/storage"
)

type CardService struct {
	storage storage.Storage
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

func (service CardService) ListCardsByLabelsAndLanguage(ctx context.Context, labels []string, language protomodels.Language) ([]protomodels.Card, error)
func (service CardService) SearchWithDictionary(ctx context.Context, cardIndex protomodels.CardIndex) (string, error)
