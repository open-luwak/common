package metadata

type EntityConfig struct {
	Entities []*Entity `toml:"entity,omitempty"`
}

type Entity struct {
	Name              string        `toml:"name"`
	LogicalDbName     string        `toml:"logical_db_name"`
	RealDbName        string        `toml:"real_db_name"`
	LogicalSchemaName string        `toml:"logical_schema_name"`
	RealSchemaName    string        `toml:"real_schema_name"`
	LogicalTableName  string        `toml:"logical_table_name"`
	RealTableName     string        `toml:"real_table_name"`
	IsView            bool          `toml:"is_view,omitempty"`
	PrimaryKey        string        `toml:"primary_key,omitempty"`
	UniqueKeys        [][]string    `toml:"unique_keys,omitempty"`
	NormalKeys        [][]string    `toml:"-"`
	ForeignKeys       []*ForeignKey `toml:"foreign_keys,omitempty"`
	Columns           []*Column     `toml:"columns,omitempty"`

	HasBeenDeleted bool `toml:"-"`
}
