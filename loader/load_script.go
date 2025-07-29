package loader

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/cast"

	"github.com/open-luwak/common/metadata"
)

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
	allowedFileExt := []string{".js", ".py", ".sql", "lua"}
	for _, v := range list {
		fileExt := filepath.Ext(v)
		if !slices.Contains(allowedFileExt, fileExt) {
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
	switch scriptType {
	case metadata.CheckpointDir:
		scriptType = metadata.CheckpointType
	case metadata.DataPreprocessDir:
		scriptType = metadata.DataPreprocessType
	case metadata.MainDir:
		scriptType = metadata.MainType
	case metadata.DatabaseUpdaterDir:
		scriptType = metadata.DatabaseUpdaterType
	case metadata.EventBroadcasterDir:
		scriptType = metadata.EventBroadcasterType
	case metadata.ResponseEnricherDir:
		scriptType = metadata.ResponseEnricherType

	// conditional role
	case metadata.ConditionalRoleDir:
		scriptType = metadata.ConditionalRoleType

	// global hooks
	case metadata.BeforeHookDir:
		scriptType = metadata.BeforeHookType
	case metadata.AfterSuccessHookDir:
		scriptType = metadata.AfterSuccessHookType
	case metadata.AfterErrorHookDir:
		scriptType = metadata.AfterErrorHookType
	case metadata.FinallyHookDir:
		scriptType = metadata.FinallyHookType

	default:
		return nil, fmt.Errorf("invalid script type: %s", scriptType)
	}

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
