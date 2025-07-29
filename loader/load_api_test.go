package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadApi(t *testing.T) {
	root := "../build/project"
	api, err := LoadApi(root)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(api)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
