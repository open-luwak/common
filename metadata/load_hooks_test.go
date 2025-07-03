package metadata

import (
	"encoding/json"
	"testing"
)

func TestLoadHooks(t *testing.T) {
	hooks, err := LoadHooks("../build/metadata")
	if err != nil {
		t.Error(err)
	}

	data, err := json.MarshalIndent(hooks, "", "  ")
	if err != nil {
		return
	}
	t.Log(string(data))
}
