package api

import (
	"github.com/open-luwak/common/metadata"
)

type Bundle struct {
	TomlBundle   *TomlBundle
	ScriptBundle *ScriptBundle

	EnableGlobalHooks bool
	HooksBundle       *HooksBundle
}

type TomlBundle struct {
	Preloader        []*metadata.DataAccessOperator
	Checkpoint       []*metadata.CheckItem
	BusinessExecutor []*metadata.DataAccessOperator
	DBUpdater        []*metadata.DataAccessOperator
	DBInserter       []*metadata.DataAccessOperator

	MetaProvider *MetaProvider
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
	Name            *ParsedName
	ConditionalRole ConditionalRoleProvider
	AutoFilter      AutoFilterProvider
	AutoPopulate    AutoPopulateProvider
}

type ConditionalRoleProvider interface {
	ApiRole(software string, method string) *[]metadata.ConditionalRole
	EntityRole(software string, entity string) *[]metadata.ConditionalRole
}

type AutoFilterProvider interface {
	ApiFilter(software string, method string) []*metadata.AutoFilter
	EntityFilter(software string, entity string) []*metadata.AutoFilter
}

type AutoPopulateProvider interface {
	ApiPopulate(software string, method string) []*metadata.AutoPopulate
	EntityPopulate(software string, entity string) []*metadata.AutoPopulate
}
