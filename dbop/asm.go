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
