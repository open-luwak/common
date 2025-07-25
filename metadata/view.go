package metadata

type ViewInfo struct {
	Name       string     `toml:"name"`
	PrimaryKey string     `toml:"primary_key"`
	UniqueKeys [][]string `toml:"unique_keys"`
	NormalKeys [][]string `toml:"-"`
}
