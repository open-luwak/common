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

func (t ResultType) String() string {
	switch t {
	case None:
		return "none"
	case MultiRows:
		return "multi_rows"
	case SingleRow:
		return "single_row"
	case AutoSingleColumn:
		return "auto_single_column"
	case SingleValue:
		return "single_value"
	case LastInsertId:
		return "last_insert_id"
	case RowsAffected:
		return "rows_affected"
	case Transaction:
		return "transaction"
	default:
		return "unknown"
	}
}

type QueryResult struct {
	ResultType     ResultType
	LastInsertId   int64
	RowsAffected   int64
	Row            map[string]any
	Rows           []map[string]any
	CheckPoint     []*QueryResult
	ResultEnricher []*QueryResult
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

func (t OperationType) String() string {
	switch t {
	case InvalidAction:
		return "invalid_action"
	case CreateOneAction:
		return "create_one_action"
	case CreateManyAction:
		return "create_many_action"
	case ReadManyAction:
		return "read_many_action"
	case ReadOneAction:
		return "read_one_action"
	case ReadValueAction:
		return "read_value_action"
	case UpdateOneAction:
		return "update_one_action"
	case UpdateManyAction:
		return "update_many_action"
	case DeleteOneAction:
		return "delete_one_action"
	case DeleteManyAction:
		return "delete_many_action"
	case TransactionAction:
		return "transaction_action"
	case DDLAction:
		return "ddl_action"
	default:
		return "unknown"
	}
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
