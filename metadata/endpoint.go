package metadata

type EndpointInfo struct {
	FullName string         `toml:"full_name"`
	Filter   map[string]any `toml:"filter"`
	Value    map[string]any `toml:"value"`
}
