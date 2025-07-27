package metadata

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cast"
)

type Storage struct {
	Env map[string]any `toml:"env,omitempty"`

	Software []*Software    `toml:"software,omitempty"`
	App      []*AppInstance `toml:"app,omitempty"`
	Endpoint []*Endpoint    `toml:"endpoint,omitempty"`
	API      []*Api         `toml:"api,omitempty"`
	Crontab  []*Job         `toml:"crontab,omitempty"`

	Db     []*DBMapping  `toml:"db,omitempty"`
	Table  []*Table      `toml:"table,omitempty"`
	View   []*Table      `toml:"view,omitempty"`
	Entity []*EntityInfo `toml:"entity,omitempty"`
}

// Load metadata information from local files, not resident in memory
//
// The kernel or plugin usually needs to be organized into its own data structure and resident in memory
func Load(dir string, subDirs []string) (*Storage, error) {
	localStorage, err := loadMetadata(dir, subDirs)
	if err != nil {
		return nil, err
	}
	// load scripts
	for _, v := range localStorage.API {
		part := strings.Split(v.Name, ".")
		if len(part) != 4 {
			continue
		}
		dbName := part[0]
		schemaName := part[1]
		tableName := part[2]
		apiName := part[3]
		baseDir := filepath.Join(dir, "api", dbName, schemaName, tableName, apiName)
		scriptInfo, err := loadScripts(baseDir)
		if err != nil {
			return nil, err
		}
		v.Scripts = scriptInfo
	}

	return localStorage, nil
}

func loadScripts(baseDir string) ([]*ScriptSource, error) {
	list, err := walkDir(baseDir)
	if err != nil {
		return nil, err
	}

	var scripts []*ScriptSource
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

	return scripts, nil
}

func parseScriptInfo(baseDir string, path string) (*ScriptSource, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid path: %s", path)
	}
	scriptType := parts[0]
	switch scriptType {
	case CheckpointDir:
		scriptType = CheckpointType
	case DataPreprocessDir:
		scriptType = DataPreprocessType
	case MainDir:
		scriptType = MainType
	case DatabaseUpdaterDir:
		scriptType = DatabaseUpdaterType
	case EventBroadcasterDir:
		scriptType = EventBroadcasterType
	case ResponseEnricherDir:
		scriptType = ResponseEnricherType

	// conditional role
	case ConditionalRoleDir:
		scriptType = ConditionalRoleType

	// global hooks
	case BeforeHookDir:
		scriptType = BeforeHookType
	case AfterSuccessHookDir:
		scriptType = AfterSuccessHookType
	case AfterErrorHookDir:
		scriptType = AfterErrorHookType
	case FinallyHookDir:
		scriptType = FinallyHookType

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

	scriptInfo := &ScriptSource{
		Name:     scriptName,
		Lang:     lang,
		Type:     scriptType,
		Code:     string(content),
		Priority: priority,
	}
	return scriptInfo, nil
}

func loadMetadata(baseDir string, subDirs []string) (*Storage, error) {
	var buf bytes.Buffer

	for _, subDir := range subDirs {
		dir := filepath.Join(baseDir, subDir)
		exists, err := dirExists(dir)
		if err != nil {
			return nil, err
		}
		if !exists {
			continue
		}

		data, err := readTomlFiles(dir)
		if err != nil {
			return nil, err
		}
		buf.Write(data)
		buf.WriteString("\n")
	}

	var localStorage = &Storage{}
	err := toml.Unmarshal(buf.Bytes(), localStorage)
	if err != nil {
		return nil, err
	}
	return localStorage, nil
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

// dirExists checks if a directory exists and is accessible
func dirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err // Return the error if it's a permission error or other problem
}
