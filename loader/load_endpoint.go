package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadEndpoint(dir string) (*metadata.EndpointConfig, error) {
	var config = &metadata.EndpointConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
