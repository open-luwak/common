package metadata

type CheckItem struct {
	Name       string `toml:"name"`
	Expression string `toml:"expression"`
}
