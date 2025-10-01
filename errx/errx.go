package errx

import (
	"fmt"
)

type Error interface {
	Code() string
	Data() any

	error
}

// New creates a new Error instance.
//
// Parameters:
//   - code: The error code.
//   - message: The error message.
//   - opts: A variadic list of optional parameters for providing additional data or an error cause.
//
// Note:
//   - `opts` can accept at most two elements.
//   - The first element can be of any type and is used to provide additional data.
//   - The second element (if provided) must be of type `error` and is used to specify the cause.
//   - If more than two elements are provided, only the first two will be used.
func New(code string, message string, opts ...any) Error {
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
	code    string
	message string
	data    any
	cause   error
}

func (e *defaultError) Error() string {
	if e.cause != nil {
		if e.message != "" {
			return fmt.Sprintf("%s, cause: %v", e.message, e.cause)
		} else {
			return e.cause.Error()
		}
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
