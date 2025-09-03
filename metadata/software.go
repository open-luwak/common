package metadata

type SoftwareConfig struct {
	SoftwareCollection []*Software `toml:"software"`
}

type Software struct {
	Name        string        `toml:"name"`
	Description string        `toml:"description"`
	AccessType  string        `toml:"access_type,omitempty"`
	Allows      []*Permission `toml:"allow,omitempty"`
	Denies      []*Permission `toml:"deny,omitempty"`
}
