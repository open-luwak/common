package metadata

import (
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func LoadAcl(dir string) (*AclConfig, error) {
	baseDir := filepath.Join(dir, "acl")
	exists, err := dirExists(baseDir)
	if err != nil {
		return nil, err
	}
	if !exists {
		return &AclConfig{}, nil
	}
	data, err := readTomlFiles(baseDir)
	if err != nil {
		return nil, err
	}
	var acl = &AclConfig{}
	err = toml.Unmarshal(data, acl)
	if err != nil {
		return nil, err
	}
	return acl, nil
}
