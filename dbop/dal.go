package dbop

import (
	"sync"
)

type ResultType int

const (
	None ResultType = iota
	MultiRows
	SingleRow
	AutoSingleColumn
	SingleValue
	LastInsertId
	RowsAffected
	Transaction
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
	InvalidAction OperationType = iota
	CreateOneAction
	CreateManyAction
	ReadManyAction
	ReadOneAction
	ReadValueAction
	UpdateOneAction
	UpdateManyAction
	DeleteOneAction
	DeleteManyAction
	TransactionAction
	DDLAction
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
