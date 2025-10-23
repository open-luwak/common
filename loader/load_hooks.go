package loader

import (
	"path/filepath"

	"github.com/open-luwak/common/metadata"
)

func LoadHooks(root string) (*metadata.HooksConfig, error) {
	dir := filepath.Join(root, defaultHooksDir)
	list, err := LoadScripts(dir)
	if err != nil {
		return nil, err
	}

	if list == nil {
		return nil, nil
	}

	hooksInfo := &metadata.HooksConfig{
		Scripts: list.Scripts,
	}
	return hooksInfo, nil
}
