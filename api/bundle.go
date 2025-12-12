package api

import (
	"github.com/open-luwak/common/metadata"
)

type Bundle struct {
	EnableGlobalHooks bool
	HooksBundle       *HooksBundle
	TomlBundle        *TomlBundle
	ScriptBundle      *ScriptBundle
	MetaProvider      *MetaProvider
}

type TomlBundle struct {
	InputValidator   []*metadata.ValidationRule
	Checks           []*metadata.CheckExpression
	Preloader        []*metadata.DataAccessOperator
	Checkpoint       []*metadata.CheckExpression
	DataPreprocessor []*metadata.DataTransformer
	BusinessExecutor []*metadata.DataAccessOperator
	DBInserter       []*metadata.DataAccessOperator
	DBUpdater        []*metadata.DataAccessOperator
	PostLoader       []*metadata.DataAccessOperator
	FieldMasker      []*metadata.FieldMasker
	ResultEnricher   []*metadata.DataTransformer
}

type ScriptBundle struct {
	InputValidator         []*metadata.ScriptSource
	Preloader              []*metadata.ScriptSource
	ConditionalRole        []*metadata.ScriptSource
	IdempotencyChecker     []*metadata.ScriptSource
	Checkpoint             []*metadata.ScriptSource
	BusinessStateValidator []*metadata.ScriptSource
	ResourceLocker         []*metadata.ScriptSource
	DataPreprocessor       []*metadata.ScriptSource
	ExternalServicePre     []*metadata.ScriptSource
	BusinessExecutor       []*metadata.ScriptSource
	DBInserter             []*metadata.ScriptSource
	DBUpdater              []*metadata.ScriptSource
	ExternalServiceInTX    []*metadata.ScriptSource
	CacheRefresher         []*metadata.ScriptSource
	AuditLogger            []*metadata.ScriptSource
	EventPublisher         []*metadata.ScriptSource
	PostLoader             []*metadata.ScriptSource
	ResultEnricher         []*metadata.ScriptSource
	FieldMasker            []*metadata.ScriptSource
	UnsupportedScripts     []*metadata.ScriptSource
}

type HooksBundle struct {
	Before             []*metadata.ScriptSource
	AfterSuccess       []*metadata.ScriptSource
	AfterError         []*metadata.ScriptSource
	Finally            []*metadata.ScriptSource
	UnsupportedScripts []*metadata.ScriptSource
}

type MetaProvider struct {
	ConditionalRole ConditionalRoleProvider
	AutoFilter      AutoFilterProvider
	AutoPopulate    AutoPopulateProvider
	Validation      ValidationProvider
}

type ConditionalRoleProvider interface {
	EntityRole(software string, entity string) []*metadata.ConditionalRole
}

type AutoFilterProvider interface {
	ApiFilter(software string, method string) []*metadata.AutoFilter
	EntityFilter(software string, entity string) []*metadata.AutoFilter
}

type AutoPopulateProvider interface {
	ApiPopulate(software string, method string) []*metadata.AutoPopulate
	EntityPopulate(software string, entity string) []*metadata.AutoPopulate
}

type ValidationProvider interface {
	ValidationRule(entity string) []*metadata.ValidationRule
	Checks(entity string) []*metadata.CheckExpression
}
