package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadAcl(t *testing.T) {
	dir := "../build/metadata/acl"
	acl, err := LoadAcl(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(acl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
