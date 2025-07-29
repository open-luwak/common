package loader

import (
	"fmt"

	"github.com/open-luwak/common/metadata"
)

func LoadGenerated(root string) (*metadata.GeneratedCache, error) {
	var config = &metadata.GeneratedConfig{}

	err := UnmarshalTomlFiles(root, config)
	if err != nil {
		return nil, err
	}

	dbMap := make(map[string]*metadata.DB, len(config.DBs))
	for _, v := range config.DBs {
		dbMap[v.Name] = v
	}

	tbMap := make(map[string]*metadata.Table, len(config.Tables)+len(config.Views))
	for _, table := range config.Tables {
		key := genMapKey(table.DbType, table)
		tbMap[key] = table
	}
	for _, view := range config.Views {
		key := genMapKey(view.DbType, view)
		tbMap[key] = view
	}

	dbTbMap := &metadata.GeneratedCache{
		DBMap: dbMap,
		TBMap: tbMap,
	}
	return dbTbMap, nil
}

func genMapKey(dbType string, v *metadata.Table) string {
	var key string
	switch dbType {
	case "postgresql", "oracle", "mssql":
		key = fmt.Sprintf("%s.%s.%s", v.DbName, v.SchemaName, v.TableName)
	default:
		key = fmt.Sprintf("%s.%s", v.DbName, v.TableName)
	}
	return key
}
