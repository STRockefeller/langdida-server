package ginserver

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/STRockefeller/langdida-server/models/protomodels"
)

func TestRelatedCardsJSONContract(t *testing.T) {
	fixtureBytes, err := os.ReadFile("testdata/related_cards.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture map[string]interface{}
	if err := json.Unmarshal(fixtureBytes, &fixture); err != nil {
		t.Fatal(err)
	}

	model := protomodels.RelatedCards{
		Index:            &protomodels.CardIndex{Name: "color", Language: 0},
		Origin:           &protomodels.CardIndex{Name: "", Language: 0},
		InOtherLanguages: []*protomodels.CardIndex{{Name: "couleur", Language: 1}},
		Synonyms:         []*protomodels.CardIndex{{Name: "hue", Language: 0}},
		Antonyms:         []*protomodels.CardIndex{{Name: "absence", Language: 0}},
		Derivatives:      []*protomodels.CardIndex{{Name: "colorful", Language: 0}},
		Others:           []*protomodels.CardIndex{{Name: "palette", Language: 0}},
	}
	actualBytes, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}
	var actual map[string]interface{}
	if err := json.Unmarshal(actualBytes, &actual); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, fixture) {
		t.Fatalf("RelatedCards JSON = %s, want fixture %s", actualBytes, fixtureBytes)
	}
	if _, exists := actual["inOtherLanguages"]; exists {
		t.Fatal("camelCase key must not be emitted")
	}
}
