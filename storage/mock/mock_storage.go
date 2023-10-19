package mock

import (
	"context"
	"time"

	"github.com/STRockefeller/go-linq"
	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/STRockefeller/langdida-server/storage"
)

type MockStorage struct {
	Cards           []protomodels.Card
	CardAssociation []protomodels.RelatedCards
}

func (storage MockStorage) ListCards(ctx context.Context, cardIndices []protomodels.CardIndex) ([]protomodels.Card, error) {
	return linq.NewLinq(storage.Cards).Where(func(c protomodels.Card) bool { return linq.NewLinq(cardIndices).Contains(*c.Index) }).ToSlice(), nil
}

func (storage MockStorage) ListCardIndices(ctx context.Context) ([]protomodels.CardIndex, error) {
	return linq.Select(storage.Cards, func(c protomodels.Card) protomodels.CardIndex { return *c.Index }).ToSlice(), nil
}

func (storage MockStorage) ListCardsWithConditions(ctx context.Context, conditions storage.ListCardsConditions) ([]protomodels.Card, error) {
	panic("not implemented") // TODO: Implement
}

func (storage MockStorage) GetAssociations(ctx context.Context, cardIndex protomodels.CardIndex) (protomodels.RelatedCards, error) {
	panic("not implemented") // TODO: Implement
}

func (storage MockStorage) CreateAssociation(ctx context.Context, conditions storage.CreateAssociationConditions) error {
	panic("not implemented") // TODO: Implement
}

// upsert to logs NewCards++
func (storage MockStorage) CreateCard(ctx context.Context, card protomodels.Card) error {
	panic("not implemented") // TODO: Implement
}

// zero values will NOT been updated
func (storage MockStorage) UpdateCard(ctx context.Context, card protomodels.Card) error {
	panic("not implemented") // TODO: Implement
}

func (storage MockStorage) DeleteCard(ctx context.Context, cardIndex protomodels.CardIndex) error {
	panic("not implemented") // TODO: Implement
}

func (storage MockStorage) GetLog(ctx context.Context, date time.Time) (protomodels.Log, error) {
	panic("not implemented") // TODO: Implement
}

func (storage MockStorage) ListLogs(ctx context.Context, from time.Time, until time.Time) ([]protomodels.Log, error) {
	panic("not implemented") // TODO: Implement
}

// upsert to logs ReviewedCards++
// update card review date
func (storage MockStorage) ReviewCard(ctx context.Context, cardIndex protomodels.CardIndex) error {
	panic("not implemented") // TODO: Implement
}
