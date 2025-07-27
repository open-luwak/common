package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadView(dir string) (*metadata.ViewConfig, error) {
	var config = &metadata.ViewConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
