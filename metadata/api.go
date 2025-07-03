package metadata

type ApiInfo struct {
	Name                string        `toml:"name"`
	Status              string        `toml:"status"`
	AccessLevel         string        `toml:"access_level"`
	ExecuteGlobalScript bool          `toml:"execute_global_script"`
	Enabled             bool          `toml:"enabled"`
	Scripts             []*ScriptInfo `toml:"-"`
}

type ScriptInfo struct {
	Name       string
	Language   string
	Type       string
	Priority   int
	SourceCode string
}
