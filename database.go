package common

import (
	"sync"
)

type AsmParam struct {
	Api           string
	Param         []any
	SessionFilter map[string]string
	SessionValues map[string]string
}

type PreparedQuery struct {
	Query     string `json:"query"`
	Bind      []any  `json:"bind"`
	QueryType string `json:"queryType"`

	CheckPoint      []*PreparedQuery `json:"checkPoint"`
	DatabaseUpdater []*PreparedQuery `json:"databaseUpdater"`
	ResultEnricher  []*PreparedQuery `json:"resultEnricher"`
}

type QueryResult struct {
	QueryType    string
	LastInsertId int64
	RowsAffected int64
	Row          map[string]any
	Rows         []map[string]any
	Total        int64
}

type DalContext struct {
	// The name of the actual physical database instance.
	PhysicalDatabaseName string

	// Container for transaction management.
	// Ensures that accessing the same database during a single API request uses the same connection,
	// and automatically commits or rolls back transactions when the request ends.
	Container *sync.Map

	TraceID  string
	ParentID string
	SpanID   string

	DebugInfo []map[string]any
}
