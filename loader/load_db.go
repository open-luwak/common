package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadDb(dir string) (*metadata.DbConfig, error) {
	var config = &metadata.DbConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
