package metadata

type AppInstance struct {
	Roles   []string       `toml:"roles,omitempty"`
	Key     string         `toml:"key"`
	Secret  string         `toml:"secret"`
	Session map[string]any `toml:"session,omitempty"`
}
