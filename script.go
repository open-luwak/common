package common

type RunnerContext struct {
	GlobalThis map[string]any
	DalCtx     *DalContext
}

type ScriptSource struct {
	EntityName string

	Name string
	Type string
	Lang string
	Code string
}
