package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestLoadDb(t *testing.T) {
	dir := "../build/generated/db"
	db, err := LoadDb(dir)
	if err != nil {
		t.Fatal(err)
	}
	data, er := toml.Marshal(db)
	if er != nil {
		t.Fatal(er)
	}
	t.Log(string(data))
}
