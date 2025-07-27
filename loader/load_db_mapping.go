package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadDbMapping(dir string) (*metadata.DBMappingConfig, error) {
	var config = &metadata.DBMappingConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
