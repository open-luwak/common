package metadata

type DataAccessOperator struct {
	Name   string `toml:"name,omitempty"`
	When   string `toml:"when,omitempty"`
	Method string `toml:"method"`
	Param  any    `toml:"param"`
	To     string `toml:"to,omitempty"`
}

type ConditionalRole struct {
	Name     string `toml:"name,omitempty"`
	Software string `toml:"software"`
	Role     string `toml:"role"`
	Check    string `toml:"check"`
}

type AutoFilter struct {
	Name     string   `toml:"name,omitempty"`
	Software []string `toml:"software"`
	Column   string   `toml:"column"`
	Operator string   `toml:"operator"`
	Value    any      `toml:"value"`
}

type AutoPopulate struct {
	Name     string   `toml:"name,omitempty"`
	Software []string `toml:"software"`
	Column   string   `toml:"column"`
	Value    any      `toml:"value"`
	On       []string `toml:"on"`
}
