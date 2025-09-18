package metadata

type AutoFilter struct {
	Software []string `toml:"software"`
	Column   string   `toml:"column"`
	Operator string   `toml:"operator"`
	Value    any      `toml:"value"`
}

type AutoPopulate struct {
	Column string `toml:"column"`
	Value  any    `toml:"value"`
}
