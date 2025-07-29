package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadTable(root string) (*metadata.TableConfig, error) {
	var config = &metadata.TableConfig{}

	dir := filepath.Join(root, defaultGeneratedDir, "table")
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
