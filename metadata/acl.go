package metadata

type Acl struct {
	Role []Role `toml:"role"`
}

type Role struct {
	Name          string  `toml:"name"`
	Description   string  `toml:"description"`
	DefaultAccess string  `toml:"default_access,omitempty"`
	Allow         []Allow `toml:"allow,omitempty"`
	Deny          []Deny  `toml:"deny,omitempty"`
}

type Allow struct {
	Resource string   `toml:"resource"`
	Action   []string `toml:"action"`
}

type Deny struct {
	Resource string   `toml:"resource"`
	Action   []string `toml:"action"`
}
