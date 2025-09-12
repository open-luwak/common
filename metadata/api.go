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

	Preloader []*Preloader `toml:"preloader,omitempty"`

	AutoFilter   []*AutoFilter   `toml:"auto_filter,omitempty"`
	AutoPopulate []*AutoPopulate `toml:"auto_populate,omitempty"`

	Validation []*ValidationRule `toml:"validation,omitempty"`
	Checking   []*CheckItem      `toml:"checking,omitempty"`

	Checkpoint []*CheckItem `toml:"checkpoint,omitempty"`

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
