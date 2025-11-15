package metadata

type Bundle struct {
	TomlBundle   *TomlBundle
	ScriptBundle *ScriptBundle

	EnableGlobalHooks bool
	HooksBundle       *HooksBundle
}

type TomlBundle struct {
	Preloader        []*DataFetcher
	Checkpoint       []*CheckItem
	BusinessExecutor []*DataAccessOperator
	DBUpdater        []*DataAccessOperator
	DBInserter       []*DataAccessOperator

	AutoFilter   []*AutoFilter
	AutoPopulate []*AutoPopulate
}

type ScriptBundle struct {
	Preloader          []*ScriptSource
	Checkpoint         []*ScriptSource
	DataPreprocessor   []*ScriptSource
	BusinessExecutor   []*ScriptSource
	ResultEnricher     []*ScriptSource
	DBUpdater          []*ScriptSource
	DBInserter         []*ScriptSource
	EventPublisher     []*ScriptSource
	ConditionalRole    []*ScriptSource
	UnsupportedScripts []*ScriptSource
}

type HooksBundle struct {
	Before             []*ScriptSource
	AfterSuccess       []*ScriptSource
	AfterError         []*ScriptSource
	Finally            []*ScriptSource
	UnsupportedScripts []*ScriptSource
}
