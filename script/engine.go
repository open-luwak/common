package script

import (
	"context"

	"github.com/open-luwak/common/bundle"
	"github.com/open-luwak/common/dbop"
	"github.com/open-luwak/common/method"
)

type RunnerContext struct {
	context.Context

	ParsedName   *method.ParsedName
	MetaProvider *bundle.MetaProvider
	DalCtx       *dbop.DalContext

	GlobalThis map[string]any
	DebugInfo  map[string]any
}
