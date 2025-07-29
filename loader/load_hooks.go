package loader

import (
	"path/filepath"
	"sort"

	"github.com/open-luwak/common/metadata"
)

func LoadHooks(root string) (*metadata.HooksConfig, error) {
	dir := filepath.Join(root, defaultHooksDir)
	list, err := LoadScripts(dir)
	if err != nil {
		return nil, err
	}

	var before, afterSuccess, afterError, finally []*metadata.ScriptSource
	for _, v := range list.Scripts {
		switch v.Type {
		case metadata.BeforeHookType:
			before = append(before, v)
		case metadata.AfterSuccessHookType:
			afterSuccess = append(afterSuccess, v)
		case metadata.AfterErrorHookType:
			afterError = append(afterError, v)
		case metadata.FinallyHookType:
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

	hooksInfo := &metadata.HooksConfig{
		Before:       before,
		AfterSuccess: afterSuccess,
		AfterError:   afterError,
		Finally:      finally,
	}
	return hooksInfo, nil
}
