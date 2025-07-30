package loader

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadDbMapping(root string) (*metadata.DBMappingConfig, error) {
	var config = &metadata.DBMappingConfig{}

	dir := filepath.Join(root, defaultDbMappingDir)
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	genDir := filepath.Join(root, defaultGeneratedDir)
	generated, err := LoadGenerated(genDir)
	if err != nil {
		return nil, err
	}

	var errs []error
	// set database version
	for _, v := range config.DBMaps {
		vv, ok := generated.DBMap[v.RealName]
		if !ok {
			errs = append(errs, fmt.Errorf("load meta/db failed: database %s not generated", v.RealName))
			continue
		}
		v.Type = vv.Type
		v.Version = vv.Version
	}

	return config, errors.Join(errs...)
}
