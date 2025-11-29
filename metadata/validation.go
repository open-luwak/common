package metadata

type ValidationRule struct {
	Field string      `toml:"field"`
	Type  string      `toml:"type,omitempty"`
	Label string      `toml:"label,omitempty"`
	Rules []*RuleItem `toml:"rules,omitempty"`
}

type RuleItem struct {
	Name      string `toml:"name,omitempty"`
	Message   string `toml:"message,omitempty"`
	Validator string `toml:"validator"`
	Value     any    `toml:"value,omitempty"`
}

type ValidationResult struct {
	Valid    bool
	Failures []RuleFailure
}

type RuleFailure struct {
	Name    string
	Field   string
	Message string
	Details map[string]any
}
