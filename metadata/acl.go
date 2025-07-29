package metadata

type AclConfig struct {
	Roles []*Role `toml:"role"`
}

type Role struct {
	Name          string   `toml:"name"`
	Description   string   `toml:"description"`
	DefaultAccess string   `toml:"default_access,omitempty"`
	Allows        []*Allow `toml:"allow,omitempty"`
	Denies        []*Deny  `toml:"deny,omitempty"`
}

type Allow struct {
	Resource string   `toml:"resource"`
	Actions  []string `toml:"action"`
}

type Deny struct {
	Resource string   `toml:"resource"`
	Actions  []string `toml:"action"`
}
