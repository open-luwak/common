package metadata

type EndpointConfig struct {
	Endpoints []*Endpoint `toml:"endpoint"`
}

type Endpoint struct {
	FullName string         `toml:"full_name"`
	Filter   map[string]any `toml:"filter"`
	Value    map[string]any `toml:"value"`
}
