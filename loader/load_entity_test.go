package loader

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml/v2"

	"github.com/open-luwak/common/metadata"
)

func TestLoadEntity(t *testing.T) {
	dir1 := "../build/generated/db"
	db1, err := LoadDb(dir1)
	if err != nil {
		t.Fatal(err)
	}
	dbMap := make(map[string]*metadata.DB, len(db1.DBs))
	for _, v := range db1.DBs {
		dbMap[v.Name] = v
	}

	dir2 := "../build/generated/table"
	table2, err := LoadTable(dir2)
	if err != nil {
		t.Fatal(err)
	}

	dir3 := "../build/generated/table"
	view3, err := LoadView(dir3)
	if err != nil {
		t.Fatal(err)
	}

	tbMap := make(map[string]*metadata.Table, len(table2.Tables)+len(view3.Views))
	for _, table := range table2.Tables {
		db2, ok := dbMap[table.DbName]
		if !ok {
			t.Errorf("database %s not found", table.DbName)
			continue
		}
		var key string
		switch db2.Type {
		case "postgresql", "oracle", "mssql":
			key = fmt.Sprintf("%s.%s.%s", table.DbName, table.SchemaName, table.TableName)
		default:
			key = fmt.Sprintf("%s.%s", table.DbName, table.TableName)
		}
		tbMap[key] = table
	}
	for _, view := range view3.Views {
		db2, ok := dbMap[view.DbName]
		if !ok {
			t.Errorf("database %s not found", view.DbName)
			continue
		}
		var key string
		switch db2.Type {
		case "postgresql", "oracle", "mssql":
			key = fmt.Sprintf("%s.%s.%s", view.DbName, view.SchemaName, view.TableName)
		default:
			key = fmt.Sprintf("%s.%s", view.DbName, view.TableName)
		}
		tbMap[key] = view
	}

	dir := "../build/metadata/db"
	entity, err := LoadEntity(dir, dbMap, tbMap)
	if err != nil {
		t.Fatal(err)
	}
	data, err := toml.Marshal(entity)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
