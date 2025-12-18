package metadata

type DatabaseNameMapping struct {
	DevName    string `toml:"dev_name"`
	DeployName string `toml:"deploy_name"`
}
