package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadDbMapping(t *testing.T) {
	root := "../build/project"
	db, err := LoadDbMapping(root)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
