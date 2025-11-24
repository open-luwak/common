package script

import (
	"github.com/open-luwak/common/api"
	"github.com/open-luwak/common/dbop"
)

type RunnerContext struct {
	ParsedName   *api.ParsedName
	MetaProvider *api.MetaProvider
	DalCtx       *dbop.DalContext

	GlobalThis map[string]any
	DebugInfo  map[string]any
}
