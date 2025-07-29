package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadDb(root string) (map[string]*metadata.DB, error) {
	var config = &metadata.DBConfig{}

	dir := filepath.Join(root, defaultGeneratedDir, "db")
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
