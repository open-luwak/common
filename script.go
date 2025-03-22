package common

type ScriptRunner interface {
	RunScript(globalThis map[string]any, pCtx PersistenceContext, sourceCode string) (any, error)
}

type GlobalThisReader interface {
	Input() any
	Env() map[string]any
	Metas() map[string]any
	Session() map[string]any
	RequestHeaders() map[string]any

	Output() any
	ResponseHeaders() map[string]any
}
