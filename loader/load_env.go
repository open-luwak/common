package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadEnv(root string) (*metadata.EnvConfig, error) {
	var env = &metadata.EnvConfig{}

	dir := filepath.Join(root, defaultEnvDir)
	err := UnmarshalTomlFiles(dir, env)
	if err != nil {
		return nil, err
	}

	return env, nil
}
