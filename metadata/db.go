package metadata

type GeneratedDbTable struct {
	DBs    []*DB    `toml:"db,omitempty"`
	Tables []*Table `toml:"table,omitempty"`
	Views  []*Table `toml:"view,omitempty"`
}

type DBConfig struct {
	DBs []*DB `toml:"db,omitempty"`
}

type DB struct {
	Name string `toml:"name"`

	Type    string `toml:"type"`
	Version string `toml:"version"`
}
