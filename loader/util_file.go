package loader

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

var (
	ErrDirNotFound      = errors.New("directory does not exist")
	ErrDirNotAccessible = errors.New("directory not accessible")
	ErrNotADirectory    = errors.New("path is not a directory")
)

func validateDir(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %v", ErrDirNotFound, path)
		}
		return fmt.Errorf("%w: %v", ErrDirNotAccessible, path)
	}
	if !fi.IsDir() {
		return fmt.Errorf("%w: %v", ErrNotADirectory, path)
	}

	return nil
}

func UnmarshalTomlFiles(dir string, v any) error {
	err := validateDir(dir)
	if err != nil {
		return err
	}

	data, err := readTomlFiles(dir)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func readTomlFiles(baseDir string) ([]byte, error) {
	var buf bytes.Buffer

	list, err := walkDir(baseDir)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		fileExt := filepath.Ext(v)
		if fileExt != ".toml" {
			continue
		}
		file := filepath.Join(baseDir, v)
		content, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		buf.Write(content)
		buf.WriteString("\n")
	}

	return buf.Bytes(), nil
}

// walkDir Returns the relative paths of all files in a directory and its subdirectories
func walkDir(rootDir string) ([]string, error) {
	var paths []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relativePath, err := filepath.Rel(rootDir, path)
			if err != nil {
				return err
			}
			paths = append(paths, relativePath)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}
	return paths, nil
}
