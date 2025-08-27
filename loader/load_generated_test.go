package loader

import (
	"encoding/json"
	"testing"
)

func TestLoadGeneratedTable(t *testing.T) {
	root := "../build/project"
	generated, err := LoadGenerated(root)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(generated, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
