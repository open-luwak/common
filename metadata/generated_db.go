package metadata

type DBConfig struct {
	DBs []*DB `toml:"db,omitempty"`
}

type DB struct {
	Name string `toml:"name"`

	Type    string `toml:"type"`
	Version string `toml:"version"`
}
