package metadata

type ValidationRule struct {
	Field string      `toml:"field"`
	Type  string      `toml:"type,omitempty"`
	Label string      `toml:"label,omitempty"`
	Rules []*RuleItem `toml:"rules,omitempty"`
}

type RuleItem struct {
	Validator      string `toml:"validator"`
	ValidatorValue []any  `toml:"validator_value,omitempty"`
	Message        string `toml:"message,omitempty"`
}
