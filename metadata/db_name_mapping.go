package metadata

type DbNameMappingConfig struct {
	DbNameMapping []DatabaseNameMapping `toml:"db_name_mapping,omitempty"`
}
type DatabaseNameMapping struct {
	DevName    string `toml:"dev_name"`
	DeployName string `toml:"deploy_name"`
}
