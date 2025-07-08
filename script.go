package common

type RunnerContext struct {
	// Provide the default entity name for getting database connections
	// only when the scripting language is native SQL.
	DefaultEntityName string

	GlobalThis map[string]any

	DalCtx *DalContext
}
