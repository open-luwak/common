package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadEntity(t *testing.T) {
	dir := "../build/project/meta"
	genDir := "../build/project/generated"

	entity, err := LoadEntity(dir, genDir)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(entity)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
