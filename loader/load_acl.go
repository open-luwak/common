package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadAcl(root string) (*metadata.AclConfig, error) {
	var config = &metadata.AclConfig{}

	dir := filepath.Join(root, defaultAclDir)
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
