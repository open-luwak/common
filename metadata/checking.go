package metadata

type CheckItem struct {
	Check   string `toml:"check"`
	Message string `toml:"message,omitempty"`
}
