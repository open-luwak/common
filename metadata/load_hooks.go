package metadata

import (
	"path/filepath"
	"sort"
)

// LoadHooks
// Deprecated
func LoadHooks(dir string) (*HooksConfig, error) {
	baseDir := filepath.Join(dir, "hooks")
	list, err := loadScripts(baseDir)
	if err != nil {
		return nil, err
	}

	var before, afterSuccess, afterError, finally []*ScriptSource
	for _, v := range list {
		switch v.Type {
		case BeforeHookType:
			before = append(before, v)
		case AfterSuccessHookType:
			afterSuccess = append(afterSuccess, v)
		case AfterErrorHookType:
			afterError = append(afterError, v)
		case FinallyHookType:
			finally = append(finally, v)
		}
	}
	sort.Slice(before, func(i, j int) bool {
		return before[i].Priority < before[j].Priority
	})
	sort.Slice(afterSuccess, func(i, j int) bool {
		return afterSuccess[i].Priority < afterSuccess[j].Priority
	})
	sort.Slice(afterError, func(i, j int) bool {
		return afterError[i].Priority < afterError[j].Priority
	})
	sort.Slice(finally, func(i, j int) bool {
		return finally[i].Priority < finally[j].Priority
	})

	hooksInfo := &HooksConfig{
		Before:       before,
		AfterSuccess: afterSuccess,
		AfterError:   afterError,
		Finally:      finally,
	}
	return hooksInfo, nil
}
