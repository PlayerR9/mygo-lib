package errors

import (
	"errors"
)

var (
	// ErrNilReceiver is an error that is returned when a receiver is nil.
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}

// ErrAfter is an error that is returned when an error occurs after an element.
type ErrAfter struct {
	// Previous is the element before the one that caused the error.
	Previous string

	// Inner is the inner error.
	Inner error
}

// Error implements the error interface.
func (e ErrAfter) Error() string {
	var prev string

	if e.Previous == "" {
		prev = "something"
	} else {
		prev = e.Previous
	}

	return "after " + prev + ": " + ErrMsgOf(e.Inner)
}

// NewErrAfter returns a new ErrAfter error.
//
// Parameters:
//   - previous: The element before the one that caused the error.
//   - inner: The inner error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"after <previous>: <inner>"
//
// Where:
//   - <previous>: The element before the one that caused the error. If empty, "something" is used.
//   - <inner>: The inner error. If nil, the default error message is used.
func NewErrAfter(previous string, inner error) error {
	return &ErrAfter{
		Previous: previous,
		Inner:    inner,
	}
}

// Unwrap returns the inner error.
//
// Returns:
//   - error: The inner error.
func (e ErrAfter) Unwrap() error {
	return e.Inner
}

// ErrNotAsExpected is an error that is returned when an element is not as expected.
type ErrNotAsExpected struct {
	// Expecteds are the expected elements.
	Expecteds []string

	// Got is the actual element.
	Got *string
}

// Error implements the error interface.
func (e ErrNotAsExpected) Error() string {
	var got string

	if e.Got == nil {
		got = "nothing"
	} else {
		got = *e.Got
	}

	var expected string

	if len(e.Expecteds) == 0 {
		expected = "nothing"
	} else {
		expected = EitherOrString(e.Expecteds)
	}

	return "expected " + expected + ", got " + got + " instead"
}

// NewErrNotAsExpected returns a new ErrNotAsExpected error.
//
// Parameters:
//   - expecteds: The expected elements.
//   - got: The actual element.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"expected <expected>, got <got> instead"
//
// Where:
//   - <expected>: The expected elements. Multiple elements use the either-or format. If empty, "nothing" is used.
//   - <got>: The actual element. If nil, "nothing" is used.
func NewErrNotAsExpected(expecteds []string, got *string) error {
	return &ErrNotAsExpected{
		Expecteds: expecteds,
		Got:       got,
	}
}
