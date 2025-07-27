package loader

import (
	"encoding/json"
	"testing"
)

func TestLoadGeneratedTable(t *testing.T) {
	dir := "../build/generated"
	generated, err := LoadGenerated(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(generated, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
