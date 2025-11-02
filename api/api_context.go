package api

import (
	"github.com/open-luwak/common/metadata"
)

var CtxKey = ContextKey{}

type ContextKey struct{}

// Context combines request and response contexts
type Context interface {
	RequestReadWriter
	ResponseReadWriter
	ServerEnvReadWriter
	AppReader
	MethodNameParser
	DebugInfoReadWriter
}

// RequestReadWriter contains request-related information
type RequestReadWriter interface {
	RequestID() string
	SetRequestID(string)

	Method() string
	SetMethod(string)

	Params() any
	SetParams(any)

	RawParams() []byte
	SetRawParams([]byte)

	Metas() map[string]any
	SetMetas(map[string]any)
}

// ResponseReadWriter contains response-related information
type ResponseReadWriter interface {
	Result() any
	SetResult(any)

	Err() error
	SetErr(error)
}

type ServerEnvReadWriter interface {
	ServerEnv() map[string]any
	SetServerEnv(map[string]any)
}

type AppReader interface {
	AppInstance() *metadata.AppInstance
}

type MethodNameParser interface {
	ParsedName() *ParsedName
}

type DebugInfoReadWriter interface {
	DebugInfo() []map[string]any
	AppendDebugInfo([]map[string]any)
}
