package metadata

type EndpointInfo struct {
	Name   string         `toml:"name"`
	Filter map[string]any `toml:"filter"`
	Value  map[string]any `toml:"value"`
}
