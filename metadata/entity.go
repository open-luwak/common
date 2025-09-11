package metadata

type EntityConfig struct {
	Entities []*Entity `toml:"entity,omitempty"`
}

type Entity struct {
	Name              string `toml:"name,omitempty"`
	LogicalDbName     string `toml:"logical_db_name,omitempty"`
	RealDbName        string `toml:"real_db_name"`
	LogicalSchemaName string `toml:"logical_schema_name,omitempty"`
	RealSchemaName    string `toml:"real_schema_name"`
	LogicalTableName  string `toml:"logical_table_name,omitempty"`
	RealTableName     string `toml:"real_table_name"`
	Label             string `toml:"label,omitempty"`
	IsView            bool   `toml:"is_view,omitempty"`

	Columns []*Column `toml:"columns,omitempty"`

	PrimaryKey string     `toml:"primary_key,omitempty"`
	UniqueKeys [][]string `toml:"unique_keys,omitempty"`
	NormalKeys [][]string `toml:"normal_keys,omitempty"`

	ForeignKeys []*ForeignKey `toml:"foreign_keys,omitempty"`

	AutoFilter      []*AutoFilter      `toml:"auto_filter,omitempty"`
	AutoPopulate    []*AutoPopulate    `toml:"auto_populate,omitempty"`
	Validation      []*ValidationRule  `toml:"validation,omitempty"`
	Checking        []*CheckItem       `toml:"checking,omitempty"`
	ConditionalRole []*ConditionalRole `toml:"conditional_role,omitempty"`

	// Deprecated
	Assignment map[string]string `toml:"assignment,omitempty"`
	// Deprecated
	Directive map[string]string `toml:"directive,omitempty"`
}
