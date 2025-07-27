package loader

import (
	"github.com/open-luwak/common/metadata"
)

func LoadDb(dir string) (map[string]*metadata.DB, error) {
	var config = &metadata.DBConfig{}

	err := UnmarshalTomlFiles(dir, config)
	if err != nil {
		return nil, err
	}

	dbMap := make(map[string]*metadata.DB, len(config.DBs))
	for _, v := range config.DBs {
		dbMap[v.Name] = v
	}

	return dbMap, nil
}
