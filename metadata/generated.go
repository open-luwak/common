package metadata

type GeneratedConfig struct {
	DBs    []*DB    `toml:"db,omitempty"`
	Tables []*Table `toml:"table,omitempty"`
	Views  []*Table `toml:"view,omitempty"`
}
