package metadata

type ApiConfig struct {
	Apis []*Api `toml:"api,omitempty"`
}

type Api struct {
	Name                string `toml:"name"`
	Status              string `toml:"status"`
	AccessLevel         string `toml:"access_level"`
	ExecuteGlobalScript bool   `toml:"execute_global_script"`
	Enabled             bool   `toml:"enabled"`

	Validation map[string]any `toml:"validation,omitempty"`
	Checking   []CheckItem    `toml:"checking,omitempty"`

	Scripts []*ScriptSource `toml:"-"`
}

type ScriptConfig struct {
	Scripts []*ScriptSource `toml:"script,omitempty"`
}

type ScriptSource struct {
	Name string
	Lang string
	Type string
	Code string

	Priority int
}
