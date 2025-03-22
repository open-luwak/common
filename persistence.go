package common

type Repository interface {
	BindNamed(query string, arg any) (string, []any, error)
	Find(query string, args ...any) (map[string]any, error)
	FindAll(query string, args ...any) ([]map[string]any, error)
	Insert(query string, args ...any) (int64, error)
	InsertBatch(query string, args ...any) (int64, error)
	// InsertNamedBatch
	// Deprecated
	InsertNamedBatch(query string, args []any) (int64, error)
	Update(query string, args ...any) (int64, error)
	Delete(query string, args ...any) (int64, error)

	Begin() error
	Commit() error
	Rollback() error

	Exec(query string, args ...any) error

	DriverName() string

	QueryTotal() int64
	SetTraceId(traceId string)
}

type PersistenceContext interface {
	GetRepository(entity string) (Repository, error)
	GetAllQueryCount() int64
	CloseAllTxn(error) error
}
