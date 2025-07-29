package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadEndpoint(root string) (*metadata.EndpointConfig, error) {
	var config = &metadata.EndpointConfig{}

	dir := filepath.Join(root, defaultEndpointDir)
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
