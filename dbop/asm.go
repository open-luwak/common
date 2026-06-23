package dbop

import (
	"github.com/open-luwak/common/metadata"
)

type SessionContext struct {
	UserID   string
	Language string
}

type AsmParam struct {
	Api          string
	Param        []any
	Session      SessionContext
	AutoFilter   []metadata.AutoFilter
	AutoPopulate []metadata.AutoPopulate
}

type PreparedQuery struct {
	QueryType string `json:"queryType"`
	Query     string `json:"query"`
	Bind      []any  `json:"bind"`

	// Query purpose, only used for subqueries like CheckPoint/ResultEnricher
	// Possible values: "uk_conflict", "fk_exists", "fk_restrict", "count", etc.
	Purpose string `json:"purpose,omitempty"`

	// Field names involved, used for generating error messages
	// For UK conflicts: unique key fields
	// For FK checks: foreign key fields
	Column []string `json:"columns,omitempty"`

	CheckPoint     []PreparedQuery `json:"checkPoint,omitempty"`
	ResultEnricher []PreparedQuery `json:"resultEnricher,omitempty"`
}
