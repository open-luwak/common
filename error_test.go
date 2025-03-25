package common

import (
	"errors"
	"fmt"
	"testing"
)

var underlyingError = fmt.Errorf("underlying error")

var testCases = []struct {
	name     string
	code     ErrorCode
	message  string
	opts     []any
	expected *Error
}{
	{
		name:     "noOpts",
		code:     "E001",
		message:  "An error occurred",
		opts:     nil,
		expected: &Error{code: "E001", message: "An error occurred", data: nil, cause: nil},
	},
	{
		name:     "withData",
		code:     "E002",
		message:  "Data error",
		opts:     []any{"data"},
		expected: &Error{code: "E002", message: "Data error", data: "data", cause: nil},
	},
	{
		name:     "withCause",
		code:     "E003",
		message:  "Error with cause",
		opts:     []any{"data", underlyingError},
		expected: &Error{code: "E003", message: "Error with cause", data: "data", cause: underlyingError},
	},
	{
		name:     "withMultipleOpts",
		code:     "E004",
		message:  "Multiple options",
		opts:     []any{"data", underlyingError, "extra"},
		expected: &Error{code: "E004", message: "Multiple options", data: "data", cause: underlyingError},
	},
}

func TestNewError(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NewError(tc.code, tc.message, tc.opts...)
			if result.code != tc.expected.code ||
				result.message != tc.expected.message ||
				result.data != tc.expected.data ||
				result.cause != tc.expected.cause {
				t.Errorf("NewError(%v, %v, %v) = %v; want %v", tc.code, tc.message, tc.opts, result, tc.expected)
			}
		})
	}
}

func TestErrorMessage(t *testing.T) {
	err := NewError("E001", "An error occurred")
	if err.Error() != "An error occurred" {
		t.Errorf("Error.Error() = %v; want %v", err, "An error occurred")
	}

	err = NewError("E003", "An error occurred", nil, underlyingError)
	if err.Error() != "An error occurred, cause: underlying error" {
		t.Errorf("Error.Unwrap() = %v; want %v", err, underlyingError)
	}
}

func TestErrorDataCause(t *testing.T) {
	err := NewError("E003", "Error with cause", "data")
	if err.Data() != "data" {
		t.Errorf("Error.Data() = %v; want %v", err, "data")
	}
	if err.Unwrap() != nil {
		t.Errorf("Unwrap did not return the expected underlying error")
	}

	err = NewError("E003", "Error with cause", underlyingError)
	if err.Data() != nil {
		t.Errorf("Error.Data() = %v; want %v", err, nil)
	}
	if err.Unwrap() != underlyingError {
		t.Errorf("Unwrap did not return the expected underlying error")
	}
}

func TestErrorUnwrap(t *testing.T) {
	err := NewError("E003", "Error with cause", "data", underlyingError)
	if !errors.Is(err, underlyingError) {
		t.Errorf("Error does not unwrap to the underlying error")
	}
	if err.Unwrap() != underlyingError {
		t.Errorf("Unwrap did not return the expected underlying error")
	}
}

func ExampleNewError() {
	err := NewError("E001", "An error occurred", "data")
	fmt.Println(err)
	fmt.Println(err.Code())
	fmt.Println(err.Data())

	err = NewError("E005", "An error occurred", underlyingError)
	fmt.Println(err)

	// Output:
	// An error occurred
	// E001
	// data
	// An error occurred, cause: underlying error
}
