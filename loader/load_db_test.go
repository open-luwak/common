package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadDb(t *testing.T) {
	dir := "../build/generated/db"
	dbMap, err := LoadDb(dir)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(dbMap)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
