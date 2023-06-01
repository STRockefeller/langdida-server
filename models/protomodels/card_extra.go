package protomodels

import (
	"encoding/json"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (l Language) MarshalJSON() ([]byte, error) {
	return json.Marshal(Language_name[int32(l)])
}

func (l *Language) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v := Language(Language_value[tmp])
	l = &v
	return nil
}

func (c Card) MarshalJSON() ([]byte, error) {
	tmp := jsonTempCard{
		Index:            c.Index,
		Labels:           c.Labels,
		Explanations:     c.Explanations,
		ExampleSentences: c.ExampleSentences,
		Familiarity:      c.Familiarity,
	}

	if c.ReviewDate == nil {
		tmp.ReviewDate = time.Now().Format("2006-01-02")
	} else {
		tmp.ReviewDate = c.ReviewDate.AsTime().Format("2006-01-02")
	}

	return json.Marshal(tmp)
}

func (c *Card) UnmarshalJSON(data []byte) error {
	var tmp jsonTempCard
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	c.Index = tmp.Index
	c.Labels = tmp.Labels
	c.Explanations = tmp.Explanations
	c.ExampleSentences = tmp.ExampleSentences
	c.Familiarity = tmp.Familiarity
	t, err := time.Parse("2006-01-02", tmp.ReviewDate)
	if err != nil {
		return err
	}
	c.ReviewDate = timestamppb.New(t)

	return nil
}

type jsonTempCard struct {
	Index            *CardIndex `json:"index,omitempty"`
	Labels           []string   `json:"labels,omitempty"`
	Explanations     []string   `json:"explanations,omitempty"`
	ExampleSentences []string   `json:"example_sentences,omitempty"`
	Familiarity      int32      `json:"familiarity,omitempty"`
	ReviewDate       string     `json:"review_date,omitempty"`
}
