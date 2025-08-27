package loader

import (
	"fmt"

	"github.com/open-luwak/common/metadata"
)

func LoadGenerated(root string) (*metadata.GeneratedCache, error) {
	dbMap, err := LoadDb(root)
	if err != nil {
		return nil, err
	}

	config, err := LoadTable(root)
	if err != nil {
		return nil, err
	}
	if config == nil {
		config = &metadata.TableConfig{}
	}

	tbMap := make(map[string]*metadata.Table, len(config.Tables))
	for _, table := range config.Tables {
		key := genMapKey(table.DbType, table)
		tbMap[key] = table
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
