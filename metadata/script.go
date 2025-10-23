package metadata

type ScriptConfig struct {
	Scripts []*ScriptSource `toml:"script,omitempty"`
}

type HooksConfig struct {
	Scripts []*ScriptSource `toml:"script,omitempty"`
}

type ScriptSource struct {
	Name string
	Lang string
	Type string
	Code string

	Priority int
}
