package script

import (
	"context"

	"github.com/open-luwak/common/bundle"
	"github.com/open-luwak/common/dbop"
	"github.com/open-luwak/common/method"
)

type RunnerContext struct {
	ctx context.Context

	ParsedName   *method.ParsedName
	MetaProvider *bundle.MetaProvider
	DalCtx       *dbop.DalContext

	GlobalThis map[string]any
	DebugInfo  map[string]any
}

func (r *RunnerContext) Context() context.Context {
	if r.ctx != nil {
		return r.ctx
	}
	return context.Background()
}

func NewRunnerContext(ctx context.Context) *RunnerContext {
	return &RunnerContext{
		ctx:        ctx,
		GlobalThis: make(map[string]any),
		DebugInfo:  make(map[string]any),
	}
}
