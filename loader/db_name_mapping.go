package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadDbNameMapping(root string) (map[string]string, error) {
	var config = &metadata.DbNameMappingConfig{}

	dir := filepath.Join(root, defaultDeploymentDir)
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, nil
	}

	mapping := make(map[string]string)
	for _, v := range config.DbNameMapping {
		mapping[v.DevName] = v.DeployName
	}

	return mapping, nil
}

func ResolveDbMapping(config *metadata.DBMappingConfig, naming map[string]string) error {
	if naming == nil {
		return nil
	}

	for _, v := range config.DBMaps {
		deployName, ok := naming[v.RealName]
		if ok {
			v.RealName = deployName
		}
	}

	return nil
}

func ResolveEntities(config *metadata.EntityConfig, naming map[string]string) error {
	if naming == nil {
		return nil
	}

	for _, v := range config.Entities {
		deployName, ok := naming[v.RealDbName]
		if ok {
			v.RealDbName = deployName
		}
	}

	return nil
}
