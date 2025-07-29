package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadView(root string) (*metadata.ViewConfig, error) {
	var config = &metadata.ViewConfig{}

	dir := filepath.Join(root, defaultGeneratedDir, "view")
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
