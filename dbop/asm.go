package dbop

type AsmParam struct {
	Api   string `json:"api"`
	Param []any  `json:"param"`

	AutoFilter   *AutoFilter   `json:"autoFilter,omitempty"`
	AutoPopulate *AutoPopulate `json:"autoPopulate,omitempty"`
}

type AutoFilter struct {
	Name     string    `json:"name,omitempty"`
	Software string    `json:"software"`
	Filters  []*Filter `json:"filters"`
}

type Filter struct {
	Column   string `json:"column"`
	Operator string `json:"operator"`
	Value    any    `json:"value,omitempty"`
}

type AutoPopulate struct {
	Name     string   `json:"name,omitempty"`
	Software string   `json:"software"`
	On       []string `json:"on"`
	Fields   []*Field `json:"fields"`
}

type Field struct {
	Column string `json:"column"`
	Value  any    `json:"value"`
}

type PreparedQuery struct {
	Query     string `json:"query"`
	Bind      []any  `json:"bind"`
	QueryType string `json:"queryType"`

	CheckPoint      []*PreparedQuery `json:"checkPoint"`
	DatabaseUpdater []*PreparedQuery `json:"databaseUpdater"`
	ResultEnricher  []*PreparedQuery `json:"resultEnricher"`
}
