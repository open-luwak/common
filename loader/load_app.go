package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadApp(root string) (*metadata.AppConfig, error) {
	var config = &metadata.AppConfig{}

	dir := filepath.Join(root, defaultAppDir)
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, err
}
