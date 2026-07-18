package review

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestReleaseWorkflowStructure(t *testing.T) {
	contents, err := os.ReadFile("../../.github/workflows/release.yml")
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		Jobs map[string]interface{} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(contents, &workflow); err != nil {
		t.Fatalf("release workflow is not valid YAML: %v", err)
	}
	for _, job := range []string{"test", "build", "release"} {
		if _, exists := workflow.Jobs[job]; !exists {
			t.Errorf("release workflow is missing %q job", job)
		}
	}
}
