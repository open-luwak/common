package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadApi(t *testing.T) {
	dir := "../build/metadata/api"
	api, err := LoadApi(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(api)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
