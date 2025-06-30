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

type EndpointMetas struct {
	Name          string
	SessionFilter map[string]any
	SessionValue  map[string]any
}

// APIContext combines request and response contexts
type APIContext interface {
	RequestReader
	ResponseReader
	RequestWriter
	ResponseWriter
	Debugger
}

// RequestReader contains request-related information
type RequestReader interface {
	RequestID() string
	RequestContext() *RequestContext
	Method() string
	Params() any
	RawParams() []byte
	Metas() map[string]any
	EndpointMetas() *EndpointMetas
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
	SetEndpointMetas(*EndpointMetas)
}

// ResponseWriter allows modification of response-related information
type ResponseWriter interface {
	SetResult(any)
	SetErr(error)
}

type Debugger interface {
	AddDebugInfo([]map[string]string)
	DebugInfo() []map[string]string
}

type Handler interface {
	Process(APIContext) error
}

type HandlerFunc func(APIContext) error

func (hf HandlerFunc) Process(ctx APIContext) error {
	return hf(ctx)
}
