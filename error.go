package common

import "fmt"

type ErrorCode string

type Error interface {
	Code() string
	Data() any

	error
}

func NewError(code ErrorCode, message string, opts ...any) Error {
	var data any
	var cause error

	if len(opts) == 1 {
		switch v := opts[0].(type) {
		case error:
			cause = v
		default:
			data = v
		}
	}
	if len(opts) > 1 {
		data = opts[0]
		cause, _ = opts[1].(error)
	}

	return &defaultError{
		code:    code,
		message: message,
		data:    data,
		cause:   cause,
	}
}

type defaultError struct {
	code    ErrorCode
	message string
	data    any
	cause   error
}

func (e *defaultError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s, cause: %v", e.message, e.cause)
	}
	return e.message
}

func (e *defaultError) Unwrap() error {
	return e.cause
}

func (e *defaultError) Code() string {
	return string(e.code)
}

func (e *defaultError) Data() any {
	return e.data
}
