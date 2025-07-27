package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadEnv(dir string) (*metadata.EnvConfig, error) {
	var env = &metadata.EnvConfig{}

	err := UnmarshalTomlFiles(dir, env)
	if err != nil {
		return nil, err
	}

	return env, nil
}
