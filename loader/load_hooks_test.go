package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadHooks(t *testing.T) {
	dir := "../build/metadata/hooks"
	hooks, err := LoadHooks(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(hooks)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
