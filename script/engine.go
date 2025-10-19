package script

import (
	"github.com/open-luwak/common/dbop"
)

type RunnerContext struct {
	// Provide the default entity name for getting database connections
	// only when the scripting language is native SQL.
	DefaultEntityName string

	GlobalThis map[string]any

	DalCtx *dbop.DalContext

	DebugInfo map[string]any
}
