package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadTable(t *testing.T) {
	root := "../build/project"
	table, err := LoadTable(root)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(table)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
