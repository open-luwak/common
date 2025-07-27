package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadTable(dir string) (*metadata.TableConfig, error) {
	var config = &metadata.TableConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
