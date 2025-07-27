package loader

import (
	"fmt"
	"log/slog"

	"github.com/open-luwak/common/metadata"
)

func LoadApi(dir string) (*metadata.ApiConfig, error) {
	var config = &metadata.ApiConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	// load scripts
	for _, v := range config.Apis {
		if err := validateName(v.Name); err != nil {
			slog.Error("loader", "error", fmt.Sprintf("invalid api name: %s", v.Name))
			continue
		}

		baseDir := scriptBaseDir(dir, v.Name)
		scriptInfo, err := LoadScripts(baseDir)
		if err != nil {
			return nil, err
		}
		v.Scripts = scriptInfo.Scripts
	}

	return config, nil
}
