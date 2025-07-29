package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadApp(t *testing.T) {
	root := "../build/project"
	app, err := LoadApp(root)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(app)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
