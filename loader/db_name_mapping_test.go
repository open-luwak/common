package loader

import (
	"encoding/json"
	"testing"
)

func TestLoadDbNameMapping(t *testing.T) {
	root := "../build/project"
	dbNameMapping, err := LoadDbNameMapping(root)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(dbNameMapping)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
