package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadDb(t *testing.T) {
	root := "../build/project"
	dbMap, err := LoadDb(root)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(dbMap)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
