package script

import (
	"github.com/open-luwak/common/api"
	"github.com/open-luwak/common/dbop"
)

type RunnerContext struct {
	// Provide the default entity name for getting database connections
	// only when the scripting language is native SQL.
	//
	// Deprecated
	DefaultEntityName string

	GlobalThis map[string]any
	DebugInfo  map[string]any

	DalCtx       *dbop.DalContext
	MetaProvider *api.MetaProvider
}
