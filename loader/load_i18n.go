package loader

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func LoadI18n(root string) (map[string]string, map[string]string, error) {
	dir := filepath.Join(root, defaultI18nDir)
	messages, err := UnmarshalI18nFiles(dir)
	if err != nil {
		return nil, nil, err
	}

	lang := make(map[string]string)
	for k, _ := range messages {
		parts := strings.SplitN(k, ".", 2)
		kk := parts[0]
		lang[kk] = kk
	}

	return lang, messages, nil
}

func UnmarshalI18nFiles(dir string) (map[string]string, error) {
	err := validateDir(dir)
	if err != nil {
		if errors.Is(err, ErrDirNotFound) {
			return nil, nil
		}
		return nil, err
	}

	data, err := readI18nFiles(dir)
	if err != nil {
		return nil, err
	}

	messages := make(map[string]string)
	var errs []error

	for k, d := range data {
		tmp := make(map[string]any)
		err = toml.Unmarshal(d, &tmp)
		if err != nil {
			return nil, err
		}
		for kk, vv := range tmp {
			key := fmt.Sprintf("%s.%s", k, kk)
			vvv, ok := vv.(string)
			if !ok {
				errs = append(errs, errors.New(key+" is not a string"))
				continue
			}
			messages[key] = vvv
		}
	}

	return messages, errors.Join(errs...)
}

func readI18nFiles(baseDir string) (map[string][]byte, error) {
	data := make(map[string][]byte)

	list, err := walkDir(baseDir)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		fileExt := filepath.Ext(v)
		if fileExt != ".toml" {
			continue
		}

		parts := strings.Split(filepath.Dir(v), string(filepath.Separator))
		key := strings.Join(parts, ".")

		file := filepath.Join(baseDir, v)
		content, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		data[key] = content
	}

	return data, nil
}
