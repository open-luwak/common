package loader

import (
	"errors"
	"path/filepath"
	"strings"
)

func validateName(name string) error {
	if name == "" {
		return errors.New("empty api name")
	}
	dotCount := strings.Count(name, ".")
	if dotCount < 3 {
		return errors.New("invalid api name")
	}
	return nil
}

func scriptBaseDir(dir string, name string) string {
	// Replace dots with path separators directly
	subDir := strings.ReplaceAll(name, ".", string(filepath.Separator))

	return filepath.Join(dir, subDir)
}
