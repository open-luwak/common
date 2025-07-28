package loader

import (
	"fmt"

	"github.com/open-luwak/common/metadata"
)

func LoadDbMapping(dir string, genDir string) (*metadata.DBMappingConfig, error) {
	var errs []error
	var config = &metadata.DBMappingConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	generated, err := LoadGenerated(genDir)
	if err != nil {
		return nil, err
	}

	// set database version
	for _, v := range config.DBMaps {
		vv, ok := generated.DBMap[v.RealName]
		if !ok {
			errs = append(errs, fmt.Errorf("database %s not generated", v.RealName))
			continue
		}
		v.Type = vv.Type
		v.Version = vv.Version
	}

	return config, nil
}
