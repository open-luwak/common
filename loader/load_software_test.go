package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadSoftware(t *testing.T) {
	root := "../build/project"
	software, err := LoadSoftware(root)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(software)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
