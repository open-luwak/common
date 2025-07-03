package metadata

type DbInfo struct {
	Type        string `toml:"type"`
	LogicalName string `toml:"logical_name"`
	RealName    string `toml:"real_name"`
	Version     string `toml:"version"`
}
