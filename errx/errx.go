package errx

import (
	"fmt"
)

type Coder interface {
	Code() string
}

type Detailer interface {
	Data() any
}

type Error interface {
	error
	Coder
	Detailer
}

type defaultError struct {
	code    string
	message string
	data    any
	cause   error
}

type Option func(*defaultError)

func WithData(data any) Option {
	return func(e *defaultError) {
		e.data = data
	}
}

func WithCause(cause error) Option {
	return func(e *defaultError) {
		e.cause = cause
	}
}

// New creates a new Error instance.
func New(code string, message string, opts ...Option) Error {
	err := &defaultError{
		code:    code,
		message: message,
	}

	for _, opt := range opts {
		opt(err)
	}

	return err
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
