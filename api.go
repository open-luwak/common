package common

import "time"

var ApiCtxKey = ContextKey{}

type ContextKey struct{}

type RequestContext struct {
	ApiName     string
	RemoteAddr  string
	RemoteHost  string
	RequestTime time.Time
}

// APIContext combines request and response contexts
type APIContext interface {
	RequestReader
	ResponseReader
	RequestWriter
	ResponseWriter
}

// RequestReader contains request-related information
type RequestReader interface {
	RequestID() string
	RequestContext() *RequestContext
	Method() string
	Params() any
	RawParams() []byte
	Metas() map[string]any
}

// ResponseReader contains response-related information
type ResponseReader interface {
	Result() any
	Err() error
}

// RequestWriter allows modification of request-related information
type RequestWriter interface {
	SetRequestID(string)
	SetMethod(string)
	SetParams(any)
	SetRawParams([]byte)
	SetMetas(map[string]any)
}

// ResponseWriter allows modification of response-related information
type ResponseWriter interface {
	SetResult(any)
	SetErr(error)
}

type Handler interface {
	Process(APIContext) error
}

type HandlerFunc func(APIContext) error

func (hf HandlerFunc) Process(ctx APIContext) error {
	return hf(ctx)
}
