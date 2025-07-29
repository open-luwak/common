package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadAcl(t *testing.T) {
	root := "../build/project"
	acl, err := LoadAcl(root)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(acl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
