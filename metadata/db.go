package metadata

type DbConfig struct {
	DBs []*DB `toml:"db,omitempty"`
}

type DB struct {
	Type    string `toml:"type"`
	Name    string `toml:"name"`
	Version string `toml:"version"`

	//Deprecated
	LogicalName string `toml:"logical_name,omitempty"`
	//Deprecated
	RealName string `toml:"real_name,omitempty"`
}
