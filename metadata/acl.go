package metadata

type AclConfig struct {
	Roles []*Role `toml:"role"`
}

type Role struct {
	Name        string        `toml:"name"`
	Description string        `toml:"description"`
	AccessType  string        `toml:"access_type,omitempty"`
	Allows      []*Permission `toml:"allow,omitempty"`
	Denies      []*Permission `toml:"deny,omitempty"`
}

type Permission struct {
	Entity string   `toml:"entity"`
	Method []string `toml:"method"`
}
