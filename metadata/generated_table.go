package metadata

type TableConfig struct {
	Tables []*Table `toml:"table,omitempty"`
}

type Table struct {
	DbType     string `toml:"db_type"`
	DbName     string `toml:"db_name"`
	SchemaName string `toml:"schema_name"`
	TableName  string `toml:"table_name"`

	IsView bool `toml:"is_view"`

	Columns []*Column `toml:"columns,omitempty"`

	PrimaryKey string     `toml:"primary_key,omitempty"`
	UniqueKeys [][]string `toml:"unique_keys,omitempty"`
	NormalKeys [][]string `toml:"normal_keys,omitempty"`

	ForeignKeys []*ForeignKey `toml:"foreign_keys,omitempty"`

	Validation []*ValidationRule `toml:"validation,omitempty"`
	Checking   []*CheckItem      `toml:"checking,omitempty"`

	// Deprecated
	Assignment map[string]any `toml:"assignment,omitempty"`
	// Deprecated
	Directive map[string]any `toml:"directive,omitempty"`
}

type Column struct {
	Name         string `toml:"name"`
	Label        string `toml:"label"`
	DataType     string `toml:"data_type"`
	Nullable     bool   `toml:"nullable"`
	HasDefault   bool   `toml:"has_default"`
	DefaultValue any    `toml:"default_value,omitempty"`
	CheckRule    string `toml:"check_rule,omitempty"`
	Comment      string `toml:"comment,omitempty"`
}

type ViewLogicalKey struct {
	DbName     string `toml:"db_name"`
	SchemaName string `toml:"schema_name"`
	TableName  string `toml:"table_name"`

	PrimaryKey string     `toml:"primary_key"`
	UniqueKeys [][]string `toml:"unique_keys"`
}

type ForeignKey struct {
	ConstraintName string      `toml:"constraint_name"`
	OnUpdate       string      `toml:"on_update"`
	OnDelete       string      `toml:"on_delete"`
	Columns        []*FKColumn `toml:"columns"`
}

type FKColumn struct {
	ColumnName            string `toml:"column_name"`
	ReferencedTableSchema string `toml:"referenced_table_schema"`
	ReferencedTableName   string `toml:"referenced_table_name"`
	ReferencedColumnName  string `toml:"referenced_column_name"`
}
