package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadApp(dir string) (*metadata.AppConfig, error) {
	var config = &metadata.AppConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, err
}
