package metadata

type Software struct {
	Name          string  `toml:"name"`
	Description   string  `toml:"description"`
	DefaultAccess string  `toml:"default_access,omitempty"`
	Allow         []Allow `toml:"allow,omitempty"`
	Deny          []Deny  `toml:"deny,omitempty"`
}
