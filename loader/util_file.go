package loader

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

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
		if errors.Is(err, ErrDirNotFound) {
			return nil
		}
		return err
	}

	data, err := catTomlFiles(dir)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func catTomlFiles(baseDir string) ([]byte, error) {
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

func readTomlFiles(baseDir string, order []string) (map[string][]byte, error) {
	fileList, err := walkDir(baseDir)
	if err != nil {
		return nil, err
	}

	fileMap := make(map[string][]string)
	for _, v := range fileList {
		if filepath.Ext(v) != ".toml" {
			continue
		}

		key := filepath.Dir(v)
		fileMap[key] = append(fileMap[key], v)
	}

	for _, list := range fileMap {
		sortByOrder(list, order)
	}

	data := make(map[string][]byte)
	for key, list := range fileMap {
		var contents [][]byte
		for _, v := range list {
			file := filepath.Join(baseDir, v)
			content, err := os.ReadFile(file)
			if err != nil {
				return nil, err
			}
			contents = append(contents, content)
		}
		data[key] = bytes.Join(contents, []byte("\n"))
	}

	return data, nil
}

func sortByOrder(slice []string, order []string) {
	orderMap := make(map[string]int)
	for i, s := range order {
		orderMap[s] = i
	}

	sort.Slice(slice, func(i, j int) bool {
		fileNameI := filepath.Base(slice[i])
		fileNameJ := filepath.Base(slice[j])

		idxI, existsI := orderMap[fileNameI]
		idxJ, existsJ := orderMap[fileNameJ]

		if existsI && existsJ {
			return idxI < idxJ
		}
		if existsI {
			return true
		}
		if existsJ {
			return false
		}
		return fileNameI < fileNameJ
	})
}
