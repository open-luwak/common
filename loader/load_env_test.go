package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadEnv(t *testing.T) {
	dir := "../build/metadata/env"
	env, err := LoadEnv(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(env)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
