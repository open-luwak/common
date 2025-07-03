package metadata

type AppInfo struct {
	Key     string         `toml:"key"`
	Secret  string         `toml:"secret"`
	Session map[string]any `toml:"session,omitempty"`
}
