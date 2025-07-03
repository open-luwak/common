package metadata

import (
	"encoding/json"
	"testing"
)

func TestLoadAcl(t *testing.T) {
	acl, err := LoadAcl("../build/metadata")
	if err != nil {
		t.Error(err)
	}

	data, err := json.MarshalIndent(acl, "", "  ")
	if err != nil {
		return
	}
	t.Log(string(data))
}
