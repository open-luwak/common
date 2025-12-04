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

	Preloader        []*DataAccessOperator `toml:"preloader,omitempty"`
	PostLoader       []*DataAccessOperator `toml:"post_loader,omitempty"`
	InputValidator   []*ValidationRule     `toml:"input_validator,omitempty"`
	Checking         []*CheckItem          `toml:"checking,omitempty"`
	Checkpoint       []*CheckItem          `toml:"checkpoint,omitempty"`
	DataPreprocessor []*DataTransformer    `toml:"data_preprocessor,omitempty"`
	ResultEnricher   []*DataTransformer    `toml:"result_enricher,omitempty"`
	BusinessExecutor []*DataAccessOperator `toml:"business_executor,omitempty"`
	DatabaseUpdater  []*DataAccessOperator `toml:"database_updater,omitempty"`
	DatabaseInserter []*DataAccessOperator `toml:"database_inserter,omitempty"`
	AutoFilter       []*AutoFilter         `toml:"auto_filter,omitempty"`
	AutoPopulate     []*AutoPopulate       `toml:"auto_populate,omitempty"`

	Scripts []*ScriptSource `toml:"-"`
}
