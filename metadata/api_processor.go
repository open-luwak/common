package metadata

type DataTransformer struct {
	Name  string `toml:"name,omitempty"`
	Field string `toml:"field,omitempty"`
	Value any    `toml:"value,omitempty"`
}

type DataAccessOperator struct {
	Name   string `toml:"name,omitempty"`
	When   string `toml:"when,omitempty"`
	To     string `toml:"to,omitempty"`
	Method string `toml:"method"`
	Param  any    `toml:"param"`
	Error  *Error `toml:"error,omitempty"`
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
	On       []string `toml:"on"`
	Column   string   `toml:"column"`
	Value    any      `toml:"value"`
}

type CheckItem struct {
	Name    string `toml:"name,omitempty"`
	Check   string `toml:"check"`
	Message string `toml:"message,omitempty"`
	Error   *Error `toml:"error,omitempty"`
}

type Error struct {
	Code    string `toml:"code"`
	Message string `toml:"message"`
}
