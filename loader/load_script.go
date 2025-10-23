package loader

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cast"

	"github.com/open-luwak/common/metadata"
)

func isAllowedFile(name string) bool {
	ext := filepath.Ext(name)
	switch ext {
	case ".js", ".py", ".sql", ".lua":
		// continue
	default:
		return false
	}

	// skip .test.*
	baseName := name[:len(name)-len(ext)]
	return !strings.HasSuffix(baseName, ".test")
}

func LoadScripts(baseDir string) (*metadata.ScriptConfig, error) {
	err := validateDir(baseDir)
	if err != nil {
		if errors.Is(err, ErrDirNotFound) {
			return nil, nil
		}
		return nil, err
	}

	list, err := walkDir(baseDir)
	if err != nil {
		return nil, err
	}

	var scripts []*metadata.ScriptSource

	for _, v := range list {
		if !isAllowedFile(v) {
			continue
		}

		info, err := parseScriptInfo(baseDir, v)
		if err != nil {
			return nil, fmt.Errorf("common.metadata: failed to parse script info from %s: %w", v, err)
		}
		scripts = append(scripts, info)
	}

	config := &metadata.ScriptConfig{
		Scripts: scripts,
	}
	return config, nil
}

func parseScriptInfo(baseDir string, path string) (*metadata.ScriptSource, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid path: %s", path)
	}
	scriptType := parts[0]

	fileName := filepath.Base(path)
	fileExt := filepath.Ext(fileName)

	fileNameParts := strings.SplitN(fileName, "_", 2)
	if len(fileNameParts) < 2 {
		return nil, fmt.Errorf("invalid path: %s", path)
	}
	priority := cast.ToInt(fileNameParts[0])

	nameAndExt := fileNameParts[1]
	scriptName := strings.TrimSuffix(nameAndExt, fileExt)

	lang := ""
	switch fileExt {
	case ".js":
		lang = "js"
	case ".py":
		lang = "python"
	case ".sql":
		lang = "sql"
	case ".lua":
		lang = "lua"

	default:
		return nil, fmt.Errorf("invalid file extension: %s", fileName)
	}

	// read source code
	file := filepath.Join(baseDir, path)
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	scriptInfo := &metadata.ScriptSource{
		Name:     scriptName,
		Lang:     lang,
		Type:     scriptType,
		Code:     string(content),
		Priority: priority,
	}
	return scriptInfo, nil
}
