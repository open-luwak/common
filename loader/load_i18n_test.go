package loader

import (
	"testing"
)

func TestLoadI18n(t *testing.T) {
	root := "../build/project"
	lang, message, err := LoadI18n(root)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(lang, message)
}
