package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadEnv(t *testing.T) {
	root := "../build/project"
	env, err := LoadEnv(root)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(env)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
