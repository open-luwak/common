package metadata

type HooksInfo struct {
	Before       []*ScriptSource
	AfterSuccess []*ScriptSource
	AfterError   []*ScriptSource
	Finally      []*ScriptSource
}
