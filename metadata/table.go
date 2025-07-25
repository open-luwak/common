package metadata

type TableInfo struct {
	Name        string        `toml:"name"`
	IsView      bool          `toml:"is_view"`
	PrimaryKey  string        `toml:"primary_key"`
	UniqueKeys  [][]string    `toml:"unique_keys"`
	ForeignKeys []*ForeignKey `toml:"foreign_keys"`
	Columns     []*RealColumn `toml:"columns"`
}

type ViewLogicalKey struct {
	Name       string     `toml:"name"`
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

type RealColumn struct {
	Name         string `toml:"name"`
	DataType     string `toml:"data_type"`
	Nullable     bool   `toml:"nullable"`
	HasDefault   bool   `toml:"has_default"`
	DefaultValue any    `toml:"default_value,omitempty"`
	CheckRule    string `toml:"check_rule,omitempty"`
	Comment      string `toml:"comment,omitempty"`
}
