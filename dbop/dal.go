package dbop

import (
	"sync"
)

type ResultType int

const (
	RetNone ResultType = iota
	RetMultiRows
	RetSingleRow
	RetSingleValue
	RetLastInsertId
	RetRowsAffected
	RetTransaction
)

type QueryResult struct {
	ResultType ResultType
	ResultName string

	Rows         []map[string]any
	Row          map[string]any
	Value        any
	LastInsertId int64
	RowsAffected int64
}

type OperationType int

const (
	OpInvalid OperationType = iota
	OpCreateOne
	OpCreateMany
	OpReadMany
	OpReadOne
	OpReadValue
	OpUpdateOne
	OpUpdateMany
	OpDeleteOne
	OpDeleteMany
	OpTransaction
	OpDDL
)

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
