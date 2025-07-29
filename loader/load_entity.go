package loader

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadEntity(root string) (*metadata.EntityConfig, error) {
	var config = &metadata.EntityConfig{}

	dir := filepath.Join(root, defaultEntityMappingDir)
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	genDir := filepath.Join(root, defaultGeneratedDir)
	generated, err := LoadGenerated(genDir)
	if err != nil {
		return nil, err
	}

	var errs []error
	// set columns, pk, uk, fk, validator
	for _, v := range config.Entities {
		vv, ok := generated.DBMap[v.RealDbName]
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
		vvv, ok := generated.TBMap[key]
		if !ok {
			if v.IsView {
				errs = append(errs, fmt.Errorf("view %s not generated", key))
			} else {
				errs = append(errs, fmt.Errorf("table %s not generated", key))
			}
			continue
		}
		v.Columns = vvv.Columns

		// view: use custom primary and unique keys
		// table: use physical table defined primary and unique keys
		if !v.IsView {
			v.PrimaryKey = vvv.PrimaryKey
			v.UniqueKeys = vvv.UniqueKeys
		}

		v.ForeignKeys = vvv.ForeignKeys

		// If no validation rules are defined, then use automatically generated validation rules.
		if len(v.Validation) == 0 {
			v.Validation = vvv.Validation
		}
		if len(v.Checking) == 0 {
			v.Checking = vvv.Checking
		}
	}

	return config, errors.Join(errs...)
}
