package metadata

type ValidationRule struct {
	Field string      `toml:"field"`
	Type  string      `toml:"type,omitempty"`
	Rules []*RuleItem `toml:"rules,omitempty"`
}

type RuleItem struct {
	Rule    string `toml:"rule"`
	Message string `toml:"message,omitempty"`
}
