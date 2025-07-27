package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadEndpoint(t *testing.T) {
	dir := "../build/metadata/endpoint"
	endpoint, err := LoadEndpoint(dir)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(endpoint)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
