package metadata

type ValidationRule struct {
	Field string      `toml:"field" json:"field"`
	Type  string      `toml:"type,omitempty" json:"type,omitempty"`
	Label string      `toml:"label,omitempty" json:"label,omitempty"`
	Rules []*RuleItem `toml:"rules,omitempty" json:"rules,omitempty"`
}

type RuleItem struct {
	Name      string `toml:"name,omitempty" json:"name,omitempty"`
	Message   string `toml:"message,omitempty" json:"message,omitempty"`
	Validator string `toml:"validator" json:"validator"`
	Value     any    `toml:"value,omitempty" json:"value,omitempty"`
}

type ValidationResult struct {
	Valid    bool           `json:"valid"`
	Failures []*RuleFailure `json:"failures,omitempty"`
}

type RuleFailure struct {
	Name    string         `json:"name,omitempty"`
	Field   string         `json:"field,omitempty"`
	Message string         `json:"message,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}

func (err *RuleFailure) Error() string {
	return err.Message
}
