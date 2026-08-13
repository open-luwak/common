package appctx

import "context"

type Context struct {
	AppRoot   string
	Namespace string
	KeyPrefix string
}

type contextKey struct{}

func With(ctx context.Context, info Context) context.Context {
	return context.WithValue(ctx, contextKey{}, info)
}

func From(ctx context.Context) (Context, bool) {
	info, ok := ctx.Value(contextKey{}).(Context)
	return info, ok
}
