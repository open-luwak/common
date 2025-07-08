package metadata

type ApiInfo struct {
	Name                string          `toml:"name"`
	Status              string          `toml:"status"`
	AccessLevel         string          `toml:"access_level"`
	ExecuteGlobalScript bool            `toml:"execute_global_script"`
	Enabled             bool            `toml:"enabled"`
	Scripts             []*ScriptSource `toml:"-"`
}

type ScriptSource struct {
	Name string
	Lang string
	Type string
	Code string

	Priority int
}
