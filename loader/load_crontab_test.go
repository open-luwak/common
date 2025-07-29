package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadCrontab(t *testing.T) {
	root := "../build/project"
	crontab, err := LoadCrontab(root)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(crontab)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
