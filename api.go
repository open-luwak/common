package common

var ApiCtxKey = ContextKey{}

type ContextKey struct{}

// APIContext combines request and response contexts
type APIContext interface {
	RequestContext
	ResponseContext
	MutableRequestContext
	MutableResponseContext
}

// RequestContext contains request-related information
type RequestContext interface {
	RequestID() string
	Method() string
	Params() any
	RawParams() []byte
	Metas() map[string]any
}

// ResponseContext contains response-related information
type ResponseContext interface {
	Result() any
	Err() error
}

// MutableRequestContext allows modification of request context
type MutableRequestContext interface {
	SetRequestID(string)
	SetMethod(string)
	SetParams(any)
	SetRawParams([]byte)
	SetMetas(map[string]any)
}

// MutableResponseContext allows modification of response context
type MutableResponseContext interface {
	SetResult(any)
	SetErr(error)
}
