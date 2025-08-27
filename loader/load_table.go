package loader

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"

	"github.com/open-luwak/common/metadata"
)

func LoadTable(root string) (*metadata.TableConfig, error) {
	dir := filepath.Join(root, defaultGeneratedDir, "table")
	config, err := UnmarshalTableFiles(dir)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func UnmarshalTableFiles(dir string) (*metadata.TableConfig, error) {
	var config = &metadata.TableConfig{}
	err := validateDir(dir)
	if err != nil {
		if errors.Is(err, ErrDirNotFound) {
			return config, nil
		}
		return nil, err
	}

	order := []string{
		"table.toml",
		"table_keys.toml",
	}

	files, err := readTomlFiles(dir, order)
	if err != nil {
		return nil, err
	}

	var errs []error

	for fileDir, fileContent := range files {
		v := &metadata.TableConfig{}
		err = toml.Unmarshal(fileContent, v)
		if err != nil {
			errs = append(errs, fmt.Errorf("error unmarshalling file %s: %w", fileDir, err))
			continue
		}
		if len(v.Tables) == 0 {
			continue
		}
		if len(v.Tables) > 1 {
			errs = append(errs, fmt.Errorf("found more than one table in %s", fileDir))
			continue
		}
		config.Tables = append(config.Tables, v.Tables...)
	}

	return config, errors.Join(errs...)
}
