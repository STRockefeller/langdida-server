package sqlite

import (
	"context"
	"errors"
	"testing"

	itime "github.com/STRockefeller/langdida-server/internal/time"
	gm "github.com/STRockefeller/langdida-server/models/gormmodels"
	"github.com/STRockefeller/langdida-server/models/protomodels"
	"github.com/golang/protobuf/ptypes/timestamp"
	"gorm.io/gorm"
)

func TestUpdateCardWritesEmptySlicesAndZeroValues(t *testing.T) {
	storage := NewStorage(":memory:", true)
	initial := gm.Card{
		Name:             "hello",
		Language:         protomodels.Language_ENGLISH,
		Labels:           gm.ArrayOfStrings{"greeting"},
		Explanations:     gm.ArrayOfStrings{"a greeting"},
		ExampleSentences: gm.ArrayOfStrings{"Hello, world."},
		Familiarity:      5,
		ReviewDate:       itime.UnixTime(12345),
	}
	if err := storage.db.Create(&initial).Error; err != nil {
		t.Fatal(err)
	}

	err := storage.UpdateCard(context.Background(), protomodels.Card{
		Index:            &protomodels.CardIndex{Name: "hello", Language: protomodels.Language_ENGLISH},
		Labels:           []string{},
		Explanations:     []string{},
		ExampleSentences: []string{},
		Familiarity:      0,
		ReviewDate:       &timestamp.Timestamp{Seconds: 0},
	})
	if err != nil {
		t.Fatal(err)
	}

	var updated gm.Card
	if err := storage.db.First(&updated, "name = ? AND language = ?", "hello", protomodels.Language_ENGLISH).Error; err != nil {
		t.Fatal(err)
	}
	if len(updated.Labels) != 0 || len(updated.Explanations) != 0 || len(updated.ExampleSentences) != 0 {
		t.Fatalf("slices were not cleared: labels=%v explanations=%v examples=%v", updated.Labels, updated.Explanations, updated.ExampleSentences)
	}
	if updated.Familiarity != 0 || updated.ReviewDate != 0 {
		t.Fatalf("zero values were not written: familiarity=%d reviewDate=%d", updated.Familiarity, updated.ReviewDate)
	}
}

func TestUpdateCardReturnsNotFound(t *testing.T) {
	storage := NewStorage(":memory:", true)
	err := storage.UpdateCard(context.Background(), protomodels.Card{
		Index:      &protomodels.CardIndex{Name: "missing", Language: protomodels.Language_ENGLISH},
		ReviewDate: &timestamp.Timestamp{},
	})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("UpdateCard error = %v, want record not found", err)
	}
}
