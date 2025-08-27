package loader

import (
	"testing"
)

func TestReadFiles(t *testing.T) {
	order := []string{
		"table.toml",
		"table_keys.toml",
	}
	files, err := readTomlFiles("../build/project/generated/table", order)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(files))
}
