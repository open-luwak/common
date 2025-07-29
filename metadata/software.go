package metadata

type SoftwareConfig struct {
	SoftwareCollection []*Software `toml:"software"`
}

type Software struct {
	Name          string   `toml:"name"`
	Description   string   `toml:"description"`
	DefaultAccess string   `toml:"default_access,omitempty"`
	Allows        []*Allow `toml:"allow,omitempty"`
	Denies        []*Deny  `toml:"deny,omitempty"`
}
