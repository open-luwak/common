package common

type DatabaseType string

type TableType string

type ColumnDataType string

type QueryType string

type PreparedQuery struct {
	Query     string    `json:"query"`
	Bind      []any     `json:"bind"`
	QueryType QueryType `json:"queryType"`

	CheckPoint      []*PreparedQuery `json:"checkPoint"`
	DatabaseUpdater []*PreparedQuery `json:"databaseUpdater"`
	ResultEnricher  []*PreparedQuery `json:"resultEnricher"`
}

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
