package metadata

type EnvConfig struct {
	Env map[string]any `toml:"env,omitempty"`
}
