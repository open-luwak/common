package metadata

type AppConfig struct {
	Apps []*AppInstance `toml:"app,omitempty"`
}

type AppInstance struct {
	Software string         `toml:"software"`
	Key      string         `toml:"key"`
	Secret   string         `toml:"secret"`
	Session  map[string]any `toml:"session,omitempty"`
}
