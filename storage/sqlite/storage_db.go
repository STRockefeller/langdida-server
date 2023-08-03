package sqlite

import (
	linq "github.com/STRockefeller/gorm-linq"
	gm "github.com/STRockefeller/langdida-server/models/gormmodels"
)

func (storage Storage) cardTable() linq.DB[gm.Card] {
	return linq.NewDB[gm.Card](storage.db)
}

func (storage Storage) logTable() linq.DB[gm.Log] {
	return linq.NewDB[gm.Log](storage.db)
}

func (storage Storage) relatedCardsTable() linq.DB[gm.RelatedCards] {
	return linq.NewDB[gm.RelatedCards](storage.db)
}
