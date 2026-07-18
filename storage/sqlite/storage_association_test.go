package sqlite

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"testing"

	gm "github.com/STRockefeller/langdida-server/models/gormmodels"
	"github.com/STRockefeller/langdida-server/models/protomodels"
	storageconditions "github.com/STRockefeller/langdida-server/storage"
)

func cardIndex(name string) protomodels.CardIndex {
	return protomodels.CardIndex{Name: name, Language: protomodels.Language_ENGLISH}
}

func createAssociation(t *testing.T, storage *Storage, from, to string, association protomodels.AssociationTypes) {
	t.Helper()
	err := storage.CreateAssociation(context.Background(), storageconditions.CreateAssociationConditions{
		CardIndex:        cardIndex(from),
		RelatedCardIndex: cardIndex(to),
		Association:      association,
	})
	if err != nil {
		t.Fatalf("CreateAssociation(%q, %q) returned error: %v", from, to, err)
	}
}

func TestCreateAssociationMergesDeduplicatesAndUpdatesBothDirections(t *testing.T) {
	storage := NewStorage(":memory:", true)

	createAssociation(t, storage, "hot", "warm", protomodels.AssociationTypes_SYNONYMS)
	createAssociation(t, storage, "hot", "cold", protomodels.AssociationTypes_ANTONYMS)
	createAssociation(t, storage, "hot", "warm", protomodels.AssociationTypes_SYNONYMS)

	hot, err := storage.GetAssociations(context.Background(), cardIndex("hot"))
	if err != nil {
		t.Fatal(err)
	}
	if len(hot.Synonyms) != 1 || hot.Synonyms[0].Name != "warm" {
		t.Fatalf("hot synonyms = %#v, want one warm association", hot.Synonyms)
	}
	if len(hot.Antonyms) != 1 || hot.Antonyms[0].Name != "cold" {
		t.Fatalf("hot antonyms = %#v, want one cold association", hot.Antonyms)
	}

	warm, err := storage.GetAssociations(context.Background(), cardIndex("warm"))
	if err != nil {
		t.Fatal(err)
	}
	if len(warm.Synonyms) != 1 || warm.Synonyms[0].Name != "hot" {
		t.Fatalf("warm synonyms = %#v, want reciprocal hot association", warm.Synonyms)
	}
	cold, err := storage.GetAssociations(context.Background(), cardIndex("cold"))
	if err != nil {
		t.Fatal(err)
	}
	if len(cold.Antonyms) != 1 || cold.Antonyms[0].Name != "hot" {
		t.Fatalf("cold antonyms = %#v, want reciprocal hot association", cold.Antonyms)
	}
}

func TestCreateAssociationRollsBackBothDirections(t *testing.T) {
	storage := NewStorage(":memory:", true)
	createAssociation(t, storage, "hot", "warm", protomodels.AssociationTypes_SYNONYMS)

	errInjected := errors.New("injected second-write failure")
	if err := storage.db.Callback().Create().Before("gorm:create").Register("test:fail_related_card", func(db *gorm.DB) {
		if card, ok := db.Statement.Dest.(*gm.RelatedCards); ok && card.Name == "fail" {
			db.AddError(errInjected)
		}
	}); err != nil {
		t.Fatal(err)
	}

	err := storage.CreateAssociation(context.Background(), storageconditions.CreateAssociationConditions{
		CardIndex:        cardIndex("hot"),
		RelatedCardIndex: cardIndex("fail"),
		Association:      protomodels.AssociationTypes_ANTONYMS,
	})
	if !errors.Is(err, errInjected) {
		t.Fatalf("CreateAssociation error = %v, want injected failure", err)
	}

	hot, err := storage.GetAssociations(context.Background(), cardIndex("hot"))
	if err != nil {
		t.Fatal(err)
	}
	if len(hot.Antonyms) != 0 {
		t.Fatalf("first write was not rolled back: antonyms = %#v", hot.Antonyms)
	}
	if len(hot.Synonyms) != 1 || hot.Synonyms[0].Name != "warm" {
		t.Fatalf("existing associations changed after rollback: %#v", hot.Synonyms)
	}
}
