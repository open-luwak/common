package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadEntity(t *testing.T) {
	root := "../build/project"
	entity, err := LoadEntity(root)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(entity)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
