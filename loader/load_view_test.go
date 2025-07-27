package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadView(t *testing.T) {
	dir := "../build/generated/view"
	view, err := LoadView(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(view)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
