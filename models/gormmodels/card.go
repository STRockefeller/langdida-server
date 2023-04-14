package gormmodels

import (
	"github.com/STRockefeller/langdida-server/internal/time"
	"github.com/STRockefeller/langdida-server/models/protomodels"
)

type Card struct {
	/* ---------------------------------- index --------------------------------- */
	Name     string               `gorm:"primaryKey"`
	Language protomodels.Language `gorm:"primaryKey"`

	Labels           ArrayOfStrings `gorm:"type:text;"`
	Explanations     ArrayOfStrings `gorm:"type:text;"`
	ExampleSentences ArrayOfStrings `gorm:"type:text;"`
	Familiarity      int32
	ReviewDate       time.UnixTime
}

func (c Card) ToProtoModel() protomodels.Card {
	return protomodels.Card{
		Index: &protomodels.CardIndex{
			Name:     c.Name,
			Language: c.Language,
		},
		Labels:           c.Labels,
		Explanations:     c.Explanations,
		ExampleSentences: c.ExampleSentences,
		Familiarity:      c.Familiarity,
		ReviewDate:       c.ReviewDate.ToTimeStamp(),
	}
}

func NewCard(c protomodels.Card) Card {
	return Card{
		Name:             c.Index.Name,
		Language:         c.Index.Language,
		Labels:           c.Labels,
		Explanations:     c.Explanations,
		ExampleSentences: c.ExampleSentences,
		Familiarity:      c.Familiarity,
		ReviewDate:       time.NewFromTimeStamp(c.ReviewDate),
	}
}
