package metadata

type DBMapping struct {
	LogicalName string `toml:"logical_name,omitempty"`
	RealName    string `toml:"real_name,omitempty"`

	Type    string `toml:"type,omitempty"`
	Version string `toml:"version,omitempty"`
}

type SchemaMapping struct {
	RealDbName  string `toml:"real_db_name"`
	LogicalName string `toml:"logical_name"`
	RealName    string `toml:"real_name"`
}
