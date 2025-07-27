package metadata

type HooksConfig struct {
	Before       []*ScriptSource
	AfterSuccess []*ScriptSource
	AfterError   []*ScriptSource
	Finally      []*ScriptSource
}
