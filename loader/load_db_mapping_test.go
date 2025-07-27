package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadDbMapping(t *testing.T) {
	dir1 := "../build/generated/db"
	dbMap, err := LoadDb(dir1)
	if err != nil {
		t.Fatal(err)
	}

	dir := "../build/metadata/db"
	db, err := LoadDbMapping(dir, dbMap)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
