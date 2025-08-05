package metadata

type ValidationRule struct {
	Field string      `toml:"field"`
	Rules []*RuleItem `toml:"rule,omitempty"`
}

type RuleItem struct {
	Rule    string `toml:"rule"`
	Message string `toml:"message,omitempty"`
}
