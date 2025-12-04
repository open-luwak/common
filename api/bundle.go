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
	Preloader        []*metadata.DataAccessOperator
	PostLoader       []*metadata.DataAccessOperator
	Checkpoint       []*metadata.CheckItem
	DataPreprocessor []*metadata.DataTransformer
	ResultEnricher   []*metadata.DataTransformer
	BusinessExecutor []*metadata.DataAccessOperator
	DBUpdater        []*metadata.DataAccessOperator
	DBInserter       []*metadata.DataAccessOperator
}

type ScriptBundle struct {
	Preloader          []*metadata.ScriptSource
	Checkpoint         []*metadata.ScriptSource
	DataPreprocessor   []*metadata.ScriptSource
	BusinessExecutor   []*metadata.ScriptSource
	ResultEnricher     []*metadata.ScriptSource
	DBUpdater          []*metadata.ScriptSource
	DBInserter         []*metadata.ScriptSource
	EventPublisher     []*metadata.ScriptSource
	ConditionalRole    []*metadata.ScriptSource
	UnsupportedScripts []*metadata.ScriptSource
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
	Checking(entity string) []*metadata.CheckItem
}
