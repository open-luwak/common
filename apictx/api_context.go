package apictx

import (
	"context"
	"net/http"

	"github.com/open-luwak/common/metadata"
	"github.com/open-luwak/common/method"
)

type ctxKey int

const (
	apiCtxKey ctxKey = iota
	requestIdKey
)

type Context struct {
	ctx context.Context

	Method      string
	RawParams   []byte
	Params      any
	Metas       map[string]any
	ServerEnv   map[string]any
	Session     map[string]any
	AppInstance *metadata.AppInstance
	ParsedName  *method.ParsedName

	RequestHeaders http.Header
	ResponseHeader http.Header

	Result    any
	Error     error
	DebugInfo []map[string]any
}

func (c *Context) Context() context.Context {
	if c.ctx != nil {
		return c.ctx
	}
	return context.Background()
}

func (c *Context) UpdateContext(ctx context.Context) {
	c.ctx = ctx
}

func New(ctx context.Context) *Context {
	return &Context{
		ctx:            ctx,
		Metas:          make(map[string]any),
		ServerEnv:      make(map[string]any),
		Session:        make(map[string]any),
		RequestHeaders: make(http.Header),
		ResponseHeader: make(http.Header),
		DebugInfo:      make([]map[string]any, 0),
	}
}

func WithApiContext(ctx context.Context, val *Context) context.Context {
	return context.WithValue(ctx, apiCtxKey, val)
}
func APIContextFrom(ctx context.Context) (*Context, bool) {
	val, ok := ctx.Value(apiCtxKey).(*Context)
	return val, ok
}

func WithRequestId(ctx context.Context, val string) context.Context {
	return context.WithValue(ctx, requestIdKey, val)
}
func RequestIdFrom(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(requestIdKey).(string)
	return val, ok
}
