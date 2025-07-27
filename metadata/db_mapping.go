package metadata

type DBMappingConfig struct {
	DBMaps []*DBMapping `toml:"db,omitempty"`
}

type DBMapping struct {
	LogicalName string `toml:"logical_name,omitempty"`
	RealName    string `toml:"real_name,omitempty"`

	Type    string `toml:"type,omitempty"`
	Version string `toml:"version,omitempty"`
}
