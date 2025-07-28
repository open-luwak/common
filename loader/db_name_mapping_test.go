package loader

import (
	"encoding/json"
	"testing"
)

func TestLoadDbNameMapping(t *testing.T) {
	dir := "../build/project/deployment"
	dbNameMapping, err := LoadDbNameMapping(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(dbNameMapping)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
