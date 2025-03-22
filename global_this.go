package common

type GlobalThisReader interface {
	Input() any
	Env() map[string]any
	Metas() map[string]any
	Session() map[string]any
	RequestHeaders() map[string]any

	Output() any
	ResponseHeaders() map[string]any
}
