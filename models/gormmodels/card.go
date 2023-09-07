package gormmodels

import (
	"encoding/json"

	"github.com/STRockefeller/go-linq"
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
	Familiarity      int32          `json:"omitempty"`
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

type RelatedCards struct {
	/* ---------------------------------- index --------------------------------- */
	Name     string               `gorm:"primaryKey"`
	Language protomodels.Language `gorm:"primaryKey"`

	Synonyms         ArrayOfStrings `gorm:"type:text;"`
	Antonyms         ArrayOfStrings `gorm:"type:text;"`
	Origin           string
	Derivatives      ArrayOfStrings `gorm:"type:text;"`
	InOtherLanguages ArrayOfStrings `gorm:"type:text;"`
	Others           ArrayOfStrings `gorm:"type:text;"`
}

func NewRelatedCards(c protomodels.RelatedCards) RelatedCards {
	var origin cardIndex
	if c.Origin != nil {
		origin = newCardIndex(*c.Origin)
	}
	return RelatedCards{
		Name:             c.Index.Name,
		Language:         c.Index.Language,
		Synonyms:         parseProtoModelCardIndexes(c.Synonyms),
		Antonyms:         parseProtoModelCardIndexes(c.Antonyms),
		Origin:           string(origin),
		Derivatives:      parseProtoModelCardIndexes(c.Derivatives),
		InOtherLanguages: parseProtoModelCardIndexes(c.InOtherLanguages),
		Others:           parseProtoModelCardIndexes(c.Others),
	}
}

func (rc RelatedCards) ToProtoModel() protomodels.RelatedCards {
	origin := cardIndex(rc.Origin).toProtoModel()
	return protomodels.RelatedCards{
		Index:            &protomodels.CardIndex{Name: rc.Name, Language: rc.Language},
		Synonyms:         toProtoModelCardIndexes(rc.Synonyms),
		Antonyms:         toProtoModelCardIndexes(rc.Antonyms),
		Origin:           &origin,
		Derivatives:      toProtoModelCardIndexes(rc.Derivatives),
		InOtherLanguages: toProtoModelCardIndexes(rc.InOtherLanguages),
		Others:           toProtoModelCardIndexes(rc.Others),
	}
}

type cardIndex string

func (c cardIndex) toProtoModel() protomodels.CardIndex {
	var cardIndex protomodels.CardIndex
	err := json.Unmarshal([]byte(c), &cardIndex)
	if err != nil {
		panic(err)
	}
	return cardIndex
}

func newCardIndex(c protomodels.CardIndex) cardIndex {
	bytes, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return cardIndex(bytes)
}

func multiCardIndexesToArrayOfString(indexes []cardIndex) ArrayOfStrings {
	return ArrayOfStrings(linq.Select(indexes, func(index cardIndex) string { return string(index) }))
}

func parseProtoModelCardIndexes(indexes []*protomodels.CardIndex) ArrayOfStrings {
	return multiCardIndexesToArrayOfString(linq.Select(indexes, func(i *protomodels.CardIndex) cardIndex {
		return newCardIndex(*i)
	}))
}

func toProtoModelCardIndexes(strings ArrayOfStrings) []*protomodels.CardIndex {
	indexes := make([]*protomodels.CardIndex, len(strings))
	for i, str := range strings {
		ptr := cardIndex(str).toProtoModel()
		indexes[i] = &ptr
	}
	return indexes
}
