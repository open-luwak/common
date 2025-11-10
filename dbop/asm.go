package dbop

import (
	"github.com/open-luwak/common/metadata"
)

type AsmParam struct {
	Api   string
	Param []any

	AutoFilter   []*metadata.AutoFilter
	AutoPopulate []*metadata.AutoPopulate
}

type PreparedQuery struct {
	QueryType string `json:"queryType"`
	Query     string `json:"query"`
	Bind      []any  `json:"bind"`

	CheckPoint      []*PreparedQuery `json:"checkPoint"`
	DatabaseUpdater []*PreparedQuery `json:"databaseUpdater"`
	ResultEnricher  []*PreparedQuery `json:"resultEnricher"`

	Result *QueryResult `json:"-"`
}

type ResultType string

const (
	None         ResultType = "none"
	MultiRows    ResultType = "multi_rows"
	SingleRow    ResultType = "single_row"
	SingleValue  ResultType = "single_value"
	LastInsertId ResultType = "last_insert_id"
	RowsAffected ResultType = "rows_affected"
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
