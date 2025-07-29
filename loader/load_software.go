package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadSoftware(root string) (*metadata.SoftwareConfig, error) {
	var config = &metadata.SoftwareConfig{}

	dir := filepath.Join(root, defaultSoftwareDir)
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
