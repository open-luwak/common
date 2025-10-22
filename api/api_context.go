package api

var CtxKey = ContextKey{}

type ContextKey struct{}

// Context combines request and response contexts
type Context interface {
	RequestReader
	ResponseReader
	RequestWriter
	ResponseWriter
	MethodNameParser
	Debugger
}

// RequestReader contains request-related information
type RequestReader interface {
	RequestID() string
	Method() string
	Params() any
	RawParams() []byte
	Metas() map[string]any

	ServerContext() map[string]any
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

	SetServerContext(map[string]any)
}

// ResponseWriter allows modification of response-related information
type ResponseWriter interface {
	SetResult(any)
	SetErr(error)
}

type MethodNameParser interface {
	ParsedName() *ParsedName
	SetParsedName(*ParsedName)
}

type Debugger interface {
	DebugInfo() []map[string]any
	AppendDebugInfo([]map[string]any)
}
