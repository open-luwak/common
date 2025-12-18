package metadata

type DB struct {
	Name string `toml:"name"`

	Type    string `toml:"type"`
	Version string `toml:"version"`
}
