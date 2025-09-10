package metadata

type CheckItem struct {
	Check   string `toml:"check"`
	Message string `toml:"message,omitempty"`
	Error   *Error `toml:"error,omitempty"`
}

type Error struct {
	Code    string `toml:"code"`
	Message string `toml:"message"`
}
