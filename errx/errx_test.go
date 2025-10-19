package errx

import (
	"errors"
	"fmt"
	"testing"
)

var underlyingError = fmt.Errorf("underlying error")

var testCases = []struct {
	name     string
	code     string
	message  string
	opts     []Option
	expected *defaultError
}{
	{
		name:     "noOpts",
		code:     "E001",
		message:  "An error occurred",
		opts:     nil,
		expected: &defaultError{code: "E001", message: "An error occurred", data: nil, cause: nil},
	},
	{
		name:     "withData",
		code:     "E002",
		message:  "Data error",
		opts:     []Option{WithData("data")},
		expected: &defaultError{code: "E002", message: "Data error", data: "data", cause: nil},
	},
	{
		name:     "withCause",
		code:     "E003",
		message:  "Error with cause",
		opts:     []Option{WithData("data"), WithCause(underlyingError)},
		expected: &defaultError{code: "E003", message: "Error with cause", data: "data", cause: underlyingError},
	},
	{
		name:     "withMultipleOpts",
		code:     "E004",
		message:  "Multiple options",
		opts:     []Option{WithData("data"), WithCause(underlyingError), WithData("extra")},
		expected: &defaultError{code: "E004", message: "Multiple options", data: "extra", cause: underlyingError},
	},
}

func TestNewError(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := New(tc.code, tc.message, tc.opts...).(*defaultError)
			if result.code != tc.expected.code ||
				result.message != tc.expected.message ||
				result.data != tc.expected.data ||
				result.cause != tc.expected.cause {
				t.Errorf("New(%v, %v, %v) = %v; want %v", tc.code, tc.message, tc.opts, result, tc.expected)
			}
		})
	}
}

func TestErrorMessage(t *testing.T) {
	err := New("E001", "An error occurred")
	if err.Error() != "An error occurred" {
		t.Errorf("Error.Error() = %v; want %v", err, "An error occurred")
	}

	err = New("E003", "An error occurred", WithData(nil), WithCause(underlyingError))
	if err.Error() != "An error occurred, cause: underlying error" {
		t.Errorf("Error.Unwrap() = %v; want %v", err, underlyingError)
	}
}

func TestErrorDataCause(t *testing.T) {
	err := New("E003", "Error with cause", WithData("data"))
	if err.Data() != "data" {
		t.Errorf("Error.Data() = %v; want %v", err, "data")
	}
	if errors.Unwrap(err) != nil {
		t.Errorf("Unwrap did not return the expected underlying error")
	}

	err = New("E003", "Error with cause", WithCause(underlyingError))
	if err.Data() != nil {
		t.Errorf("Error.Data() = %v; want %v", err, nil)
	}
	if errors.Unwrap(err) != underlyingError {
		t.Errorf("Unwrap did not return the expected underlying error")
	}
}

func TestErrorUnwrap(t *testing.T) {
	err := New("E003", "Error with cause", WithData("data"), WithCause(underlyingError))
	if !errors.Is(err, underlyingError) {
		t.Errorf("Error does not unwrap to the underlying error")
	}
	if errors.Unwrap(err) != underlyingError {
		t.Errorf("Unwrap did not return the expected underlying error")
	}
}

func ExampleNew() {
	err := New("E001", "An error occurred", WithData("data"))
	fmt.Println(err)
	fmt.Println(err.Code())
	fmt.Println(err.Data())

	err = New("E005", "An error occurred", WithCause(underlyingError))
	fmt.Println(err)

	// Output:
	// An error occurred
	// E001
	// data
	// An error occurred, cause: underlying error
}
