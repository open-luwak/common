package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadDbMapping(t *testing.T) {
	genDir := "../build/generated/db"
	dir := "../build/metadata/db"
	db, err := LoadDbMapping(dir, genDir)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
