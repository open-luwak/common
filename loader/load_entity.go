package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadEntity(dir string) (*metadata.EntityConfig, error) {
	var config = &metadata.EntityConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
