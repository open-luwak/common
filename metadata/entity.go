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
