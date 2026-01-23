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

type DefaultError struct {
	code    string
	message string
	data    any
	cause   error
}

type Option func(*DefaultError)

func WithData(data any) Option {
	return func(e *DefaultError) {
		e.data = data
	}
}

func WithCause(cause error) Option {
	return func(e *DefaultError) {
		e.cause = cause
	}
}

// New creates a new Error instance.
func New(code string, message string, opts ...Option) *DefaultError {
	err := &DefaultError{
		code:    code,
		message: message,
	}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

func (e *DefaultError) Error() string {
	if e.cause != nil {
		if e.message != "" {
			return fmt.Sprintf("%s, cause: %v", e.message, e.cause)
		} else {
			return e.cause.Error()
		}
	}
	return e.message
}

func (e *DefaultError) Unwrap() error {
	return e.cause
}

func (e *DefaultError) Code() string {
	return string(e.code)
}

func (e *DefaultError) Data() any {
	return e.data
}

func (e *DefaultError) SetData(data any) {
	e.data = data
}
func (e *DefaultError) SetCause(cause error) {
	e.cause = cause
}
