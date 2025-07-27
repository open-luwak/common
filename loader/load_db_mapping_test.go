package loader

import (
	"testing"

	"github.com/pelletier/go-toml/v2"

	"github.com/open-luwak/common/metadata"
)

func TestLoadDbMapping(t *testing.T) {
	dir1 := "../build/generated/db"
	db1, err := LoadDb(dir1)
	if err != nil {
		t.Fatal(err)
	}
	dbMap := make(map[string]*metadata.DB, len(db1.DBs))
	for _, v := range db1.DBs {
		dbMap[v.Name] = v
	}

	dir := "../build/metadata/db"
	db, err := LoadDbMapping(dir, dbMap)
	if err != nil {
		t.Fatal(err)
	}

	data, err := toml.Marshal(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
