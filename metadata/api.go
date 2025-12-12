package metadata

type ApiConfig struct {
	Apis []*Api `toml:"api,omitempty"`
}

type Api struct {
	Name                string `toml:"name"`
	Status              string `toml:"status"`
	AccessLevel         string `toml:"access_level"`
	ExecuteGlobalScript bool   `toml:"execute_global_script"`
	Enabled             bool   `toml:"enabled"`

	InputValidator   []*ValidationRule     `toml:"input_validator,omitempty"`
	Checks           []*CheckExpression    `toml:"checks,omitempty"`
	Preloader        []*DataAccessOperator `toml:"pre_loader,omitempty"`
	Checkpoint       []*CheckExpression    `toml:"checkpoint,omitempty"`
	AutoFilter       []*AutoFilter         `toml:"auto_filter,omitempty"`
	AutoPopulate     []*AutoPopulate       `toml:"auto_populate,omitempty"`
	DataPreprocessor []*DataTransformer    `toml:"data_preprocessor,omitempty"`
	BusinessExecutor []*DataAccessOperator `toml:"business_executor,omitempty"`
	DatabaseInserter []*DataAccessOperator `toml:"database_inserter,omitempty"`
	DatabaseUpdater  []*DataAccessOperator `toml:"database_updater,omitempty"`
	PostLoader       []*DataAccessOperator `toml:"post_loader,omitempty"`
	FieldMasker      []*FieldMasker        `toml:"field_masker,omitempty"`
	ResultEnricher   []*DataTransformer    `toml:"result_enricher,omitempty"`

	Scripts []*ScriptSource `toml:"-"`
}
