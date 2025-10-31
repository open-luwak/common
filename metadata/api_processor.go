package metadata

type DataFetcher struct {
	Name   string `toml:"name,omitempty"`
	Method string `toml:"method"`
	Param  any    `toml:"param"`
	To     string `toml:"to"`
}

type DataAccessOperator struct {
	Name   string `toml:"name,omitempty"`
	Method string `toml:"method"`
	Param  any    `toml:"param"`
}

type ConditionalRole struct {
	Name     string `toml:"name,omitempty"`
	Software string `toml:"software"`
	Role     string `toml:"role"`
	Check    string `toml:"check"`
}

type AutoFilter struct {
	Name     string    `toml:"name,omitempty"`
	Software []string  `toml:"software"`
	Filters  []*Filter `toml:"filters"`
}

type Filter struct {
	Column   string `toml:"column"`
	Operator string `toml:"operator"`
	Value    any    `toml:"value"`
}

type AutoPopulate struct {
	Name     string   `toml:"name,omitempty"`
	Software []string `toml:"software"`
	On       []string `toml:"on"`
	Fields   []*Field `toml:"fields"`
}

type Field struct {
	Column string `toml:"column"`
	Value  any    `toml:"value"`
}
