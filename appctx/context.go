package appctx

import "context"

type Context struct {
	appRoot   string
	namespace string
	keyPrefix string
}

func New(appRoot, namespace, keyPrefix string) *Context {
	return &Context{appRoot, namespace, keyPrefix}
}

func (c *Context) AppRoot() string   { return c.appRoot }
func (c *Context) Namespace() string { return c.namespace }
func (c *Context) KeyPrefix() string { return c.keyPrefix }

type contextKey struct{}

func With(ctx context.Context, cfg *Context) context.Context {
	return context.WithValue(ctx, contextKey{}, cfg)
}

func From(ctx context.Context) (*Context, bool) {
	cfg, ok := ctx.Value(contextKey{}).(*Context)
	return cfg, ok
}
