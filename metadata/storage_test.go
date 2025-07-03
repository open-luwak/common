package metadata

import (
	"encoding/json"
	"testing"
)

func TestLoad(t *testing.T) {
	baseDir := "../build/metadata"
	subDirs := []string{"env", "app", "db", "api", "endpoint", "crontab"}
	localStorage, err := Load(baseDir, subDirs)
	if err != nil {
		t.Error(err)
	}

	data, err := json.MarshalIndent(localStorage, "", "  ")
	if err != nil {
		return
	}
	t.Log(string(data))
}
