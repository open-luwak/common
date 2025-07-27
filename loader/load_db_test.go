package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadDb(t *testing.T) {
	dir := "../build/metadata/db"
	db, err := LoadDb(dir)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
