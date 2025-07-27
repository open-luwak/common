package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadSoftware(dir string) (*metadata.SoftwareConfig, error) {
	var config = &metadata.SoftwareConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
