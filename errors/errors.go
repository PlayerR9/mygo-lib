package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrNilReceiver is an error that is returned when a receiver is nil.
	//
	// Format:
	//   "receiver must not be nil"
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}

// ErrNotAsExpected is an error that is returned when an element is not as expected.
type ErrNotAsExpected[T any] struct {
	// Quote indicates whether the expected elements should be quoted.
	Quote bool

	// Kind is the kind of the expected elements.
	Kind string

	// Expecteds is the expected elements.
	Expecteds []T

	// Got is the actual element.
	Got any
}

// Error implements the error interface.
func (e ErrNotAsExpected[T]) Error() string {
	var expected string

	if len(e.Expecteds) == 0 {
		expected = "something"
	} else {
		elems := make([]string, 0, len(e.Expecteds))

		for _, elem := range e.Expecteds {
			elems = append(elems, fmt.Sprint(elem))
		}

		if e.Quote {
			for i := 0; i < len(elems); i++ {
				elems[i] = strconv.Quote(elems[i])
			}
		}

		expected = EitherOrString(elems)
	}

	var got string

	if e.Got == nil {
		got = "nothing"
	} else if e.Quote {
		got = strconv.Quote(fmt.Sprint(e.Got))
	} else {
		got = fmt.Sprint(e.Got)
	}

	var builder strings.Builder

	builder.WriteString("expected ")

	if e.Kind != "" {
		builder.WriteString(e.Kind)
		builder.WriteString(" to be ")
	}

	builder.WriteString(expected)
	builder.WriteString(", got ")
	builder.WriteString(got)

	return builder.String()
}

// NewErrNotAsExpected returns a new ErrNotAsExpected error.
//
// Parameters:
//   - quote: Indicates whether the expected elements should be quoted.
//   - kind: The kind of the expected elements. If empty, it will be omitted.
//   - got: The actual element.
//   - expecteds: The expected elements.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	expected <kind> to be <expecteds>, got <got>
func NewErrNotAsExpected[T any](quote bool, kind string, got any, expecteds ...T) error {
	return &ErrNotAsExpected[T]{
		Quote:     quote,
		Kind:      kind,
		Expecteds: expecteds,
		Got:       got,
	}
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

// ErrPanic is an error that is returned when a panic occurs.
type ErrPanic struct {
	// value is the value that was passed to panic.
	value any
}

// Error implements the error interface.
func (e ErrPanic) Error() string {
	return fmt.Sprintf("panic: %v", e.value)
}

// NewErrPanic creates a new ErrPanic error.
//
// Parameters:
//   - value: The value that was passed to panic.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"panic: <value>"
//
// Where:
//   - <value>: The value that was passed to panic.
func NewErrPanic(value any) error {
	return &ErrPanic{
		value: value,
	}
}
