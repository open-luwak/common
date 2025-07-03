package metadata

type EntityInfo struct {
	Name              string        `toml:"name"`
	LogicalDbName     string        `toml:"logical_db_name"`
	RealDbName        string        `toml:"real_db_name"`
	LogicalSchemaName string        `toml:"logical_schema_name"`
	RealSchemaName    string        `toml:"real_schema_name"`
	LogicalTableName  string        `toml:"logical_table_name"`
	RealTableName     string        `toml:"real_table_name"`
	IsView            bool          `toml:"is_view"`
	PrimaryKey        string        `toml:"primary_key"`
	UniqueKeys        [][]string    `toml:"unique_keys"`
	NormalKeys        [][]string    `toml:"-"`
	ForeignKeys       []*ForeignKey `toml:"foreign_keys"`
	Columns           []*RealColumn `toml:"columns"`

	HasBeenDeleted bool `toml:"-"`
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
	DefaultValue string `toml:"default_value"`
	CheckRule    string `toml:"check_rule"`
	Comment      string `toml:"comment,omitempty"`
}
