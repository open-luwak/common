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
	RequestIdKey
	TenantIdKey
	AppIdKey
)

type Context struct {
	context.Context

	RequestID   string
	TenantID    string
	AppID       string
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

func (c *Context) Value(key any) any {
	switch key {
	case RequestIdKey:
		return c.RequestID
	case TenantIdKey:
		return c.TenantID
	case AppIdKey:
		return c.AppID
	default:
		return c.Context.Value(key)
	}
}

func New(ctx context.Context) *Context {
	return &Context{
		Context:        ctx,
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

func FromApiContext(ctx context.Context) (*Context, bool) {
	val, ok := ctx.Value(apiCtxKey).(*Context)
	return val, ok
}

func GetTenantId(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(TenantIdKey).(string)
	return val, ok
}

func GetAppId(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(AppIdKey).(string)
	return val, ok
}
