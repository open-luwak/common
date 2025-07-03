package metadata

type HooksInfo struct {
	Before       []*ScriptInfo
	AfterSuccess []*ScriptInfo
	AfterError   []*ScriptInfo
	Finally      []*ScriptInfo
}
