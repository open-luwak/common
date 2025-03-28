package common

import (
	"sync"
)

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

type QueryResult struct {
	Row          map[string]any
	Rows         []map[string]any
	RowsAffected int64
	LastInsertId int64
	Total        int64
}

type DalContext struct {
	DefaultEntity string
	Container     sync.Map

	TraceID  string
	ParentID string
	SpanID   string
}
