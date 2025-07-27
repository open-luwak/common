package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadAcl(dir string) (*metadata.AclConfig, error) {
	var config = &metadata.AclConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
