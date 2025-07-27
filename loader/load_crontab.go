package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadCrontab(dir string) (*metadata.CrontabConfig, error) {
	var config = &metadata.CrontabConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, err
}
