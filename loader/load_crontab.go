package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadCrontab(root string) (*metadata.CrontabConfig, error) {
	var config = &metadata.CrontabConfig{}

	dir := filepath.Join(root, defaultCrontabDir)
	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, err
}
