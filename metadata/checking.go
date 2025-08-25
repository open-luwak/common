package metadata

type CheckItem struct {
	Expression string `toml:"expression"`
	Message    string `toml:"message,omitempty"`
}
