package dbop

import (
	"sync"
)

type QueryResult struct {
	QueryType    string
	LastInsertId int64
	RowsAffected int64
	Row          map[string]any
	Rows         []map[string]any
	Total        int64
}

type DalContext struct {
	// Container for transaction management.
	// Ensures that accessing the same database during a single API request uses the same connection,
	// and automatically commits or rolls back transactions when the request ends.
	Container *sync.Map

	TraceID  string
	ParentID string
	SpanID   string

	DebugInfo []map[string]any
}
