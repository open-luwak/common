package loader

import (
	"fmt"

	"github.com/open-luwak/common/metadata"
)

func LoadEntity(dir string, dbMap map[string]*metadata.DB, tbMap map[string]*metadata.Table) (*metadata.EntityConfig, error) {
	var errs []error
	var config = &metadata.EntityConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	// set columns, pk, uk, fk, validator
	for _, v := range config.Entities {
		vv, ok := dbMap[v.RealDbName]
		if !ok {
			errs = append(errs, fmt.Errorf("database %s not generated", v.RealDbName))
			continue
		}
		var key string
		switch vv.Type {
		case "postgresql", "oracle", "mssql":
			key = fmt.Sprintf("%s.%s.%s", v.RealDbName, v.RealSchemaName, v.RealTableName)
		default:
			key = fmt.Sprintf("%s.%s", v.RealDbName, v.RealTableName)
		}
		vvv, ok := tbMap[key]
		if !ok {
			if v.IsView {
				errs = append(errs, fmt.Errorf("view %s not generated", key))
			} else {
				errs = append(errs, fmt.Errorf("table %s not generated", key))
			}
			continue
		}
		v.Columns = vvv.Columns
		v.PrimaryKey = vvv.PrimaryKey
		v.UniqueKeys = vvv.UniqueKeys
		v.ForeignKeys = vvv.ForeignKeys
	}

	return config, nil
}
