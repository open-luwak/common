package loader

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"

	"github.com/open-luwak/common/metadata"
)

func LoadEntity(root string) (*metadata.EntityConfig, error) {
	dir := filepath.Join(root, defaultEntityMappingDir)
	config, err := UnmarshalEntityFiles(dir)
	if err != nil {
		return nil, err
	}

	generated, err := LoadGenerated(root)
	if err != nil {
		return nil, err
	}

	var errs []error
	// set columns, pk, uk, fk, validator
	for _, v := range config.Entities {
		vv, ok := generated.DBMap[v.RealDbName]
		if !ok {
			errs = append(errs, fmt.Errorf("load %s.%s.%s failed: database %s not generated", v.LogicalDbName, v.LogicalSchemaName, v.LogicalTableName, v.RealDbName))
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
				errs = append(errs, fmt.Errorf("load %s.%s.%s failed: view %s not generated", v.LogicalDbName, v.LogicalSchemaName, v.LogicalTableName, key))
			} else {
				errs = append(errs, fmt.Errorf("load %s.%s.%s failed: table %s not generated", v.LogicalDbName, v.LogicalSchemaName, v.LogicalTableName, key))
			}
			continue
		}
		v.Columns = vvv.Columns

		if v.Label == "" {
			v.Label = vvv.Label
		}

		// view: use custom primary and unique keys
		// table: use physical table defined primary and unique keys
		if !v.IsView {
			v.PrimaryKey = vvv.PrimaryKey
			v.UniqueKeys = vvv.UniqueKeys
		}

		v.ForeignKeys = vvv.ForeignKeys

		v.AutoFilter = vvv.AutoFilter
		v.AutoPopulate = vvv.AutoPopulate

		v.Validation = mergeValidation(v.Validation, vvv.Validation)

		if len(v.Checks) == 0 {
			v.Checks = vvv.Checks
		}

		v.ConditionalRole = vvv.ConditionalRole
	}

	return config, errors.Join(errs...)
}

func mergeValidation(entityValidation, tableValidation []*metadata.ValidationRule) []*metadata.ValidationRule {
	var out []*metadata.ValidationRule

	hasRule := make(map[string]int)

	genKey := func(field, rule string) string {
		return fmt.Sprintf("%s.%s", field, rule)
	}

	for _, v := range tableValidation {
		for i, vv := range v.Rules {
			key := genKey(v.Field, vv.Validator)
			hasRule[key] = i
		}
		out = append(out, v)
	}

	// If no validation rules are defined, then use automatically generated validation rules.
	if len(entityValidation) == 0 {
		return out
	}

	for _, v := range out {
		for _, e := range entityValidation {
			if e.Field != v.Field {
				continue
			}

			if e.Type != "" {
				v.Type = e.Type
			}
			if e.Label != "" {
				v.Label = e.Label
			}
			for _, rule := range e.Rules {
				key := genKey(v.Field, rule.Validator)
				if i, ok := hasRule[key]; ok {
					v.Rules[i] = rule // 覆盖
				} else {
					v.Rules = append(v.Rules, rule) // 追加
				}
			}
		}
	}

	return out
}

func UnmarshalEntityFiles(dir string) (*metadata.EntityConfig, error) {
	var config = &metadata.EntityConfig{}

	err := validateDir(dir)
	if err != nil {
		if errors.Is(err, ErrDirNotFound) {
			return config, nil
		}
		return nil, err
	}

	order := []string{
		"entity.toml",
		"entity_keys.toml",
	}

	files, err := readTomlFiles(dir, order)
	if err != nil {
		return nil, err
	}

	var errs []error

	for fileDir, fileContent := range files {
		v := &metadata.EntityConfig{}
		err = toml.Unmarshal(fileContent, v)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if len(v.Entities) == 0 {
			continue
		}
		if len(v.Entities) > 1 {
			errs = append(errs, fmt.Errorf("load entity failed: more than one entity in %s", fileDir))
			continue
		}

		parts := strings.Split(fileDir, string(filepath.Separator))
		if len(parts) != 3 {
			errs = append(errs, fmt.Errorf("load entity failed: must be a three-level directory: %s", fileDir))
			continue
		}

		entity := &metadata.Entity{
			Name:              strings.Join(parts, "."),
			LogicalDbName:     parts[0],
			RealDbName:        v.Entities[0].RealDbName,
			LogicalSchemaName: parts[1],
			RealSchemaName:    v.Entities[0].RealSchemaName,
			LogicalTableName:  parts[2],
			RealTableName:     v.Entities[0].RealTableName,
			IsView:            v.Entities[0].IsView,
			//If it's a view, it may contain custom primary keys and unique keys.
			PrimaryKey: v.Entities[0].PrimaryKey,
			UniqueKeys: v.Entities[0].UniqueKeys,
			Validation: v.Entities[0].Validation,
			Checks:     v.Entities[0].Checks,
		}

		config.Entities = append(config.Entities, entity)
	}

	return config, errors.Join(errs...)
}
