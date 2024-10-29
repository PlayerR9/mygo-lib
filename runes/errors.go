package runes

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/PlayerR9/mygo-lib/common"
)

var (
	// ErrBadEncoding occurs when an input (slice of bytes or string) is not valid utf-8.
	// This error can be checked using the == operator.
	//
	// Format:
	// 	"invalid utf-8"
	ErrBadEncoding error
)

func init() {
	ErrBadEncoding = errors.New("invalid utf-8")
}

// ErrAt occurs when an error occurs at a specific index.
type ErrAt struct {
	// Idx is the index at which the error occurred.
	Idx int

	// Inner is the inner error.
	Inner error
}

// Error implements the error interface.
func (e ErrAt) Error() string {
	var reason string

	if e.Inner == nil {
		reason = "something went wrong"
	} else {
		reason = e.Inner.Error()
	}

	return fmt.Sprintf("at index %d: %s", e.Idx, reason)
}

// NewErrAt returns a new ErrAt from the given index and inner error.
//
// Parameters:
//   - idx: The index at which the error occurred.
//   - inner: The inner error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"at index <idx>: <reason>"
//
// Where:
//   - <idx>: The index at which the error occurred.
//   - <reason>: The reason for the error. If nil, "something went wrong" is used instead.
func NewErrAt(idx int, inner error) error {
	return &ErrAt{
		Idx:   idx,
		Inner: inner,
	}
}

// Unwrap implements the errors.Wrapper interface.
//
// Returns:
//   - error: The inner error.
func (e ErrAt) Unwrap() error {
	return e.Inner
}

// ErrNotAsExpected occurs when a rune is not as expected.
type ErrNotAsExpected struct {
	// Quote if true, the runes will be quoted before being printed.
	Quote bool

	// Kind is the kind of the rune that is not as expected.
	Kind string

	// Expecteds are the runes that were expected.
	Expecteds []rune

	// Got is the actual rune.
	Got *rune
}

// Error implements the error interface.
func (e ErrNotAsExpected) Error() string {
	var kind string

	if e.Kind != "" {
		kind = e.Kind + " to be "
	}

	var got string

	if e.Got == nil {
		got = "nothing"
	} else if e.Quote {
		got = strconv.QuoteRune(*e.Got)
	} else {
		got = string(*e.Got)
	}

	var builder strings.Builder

	builder.WriteString("expected ")
	builder.WriteString(kind)

	elems := make([]string, 0, len(e.Expecteds))

	if e.Quote {
		for _, elem := range e.Expecteds {
			str := strconv.QuoteRune(elem)
			elems = append(elems, str)
		}
	} else {
		for _, elem := range e.Expecteds {
			str := string(elem)
			elems = append(elems, str)
		}
	}

	builder.WriteString(common.EitherOrString(elems))
	builder.WriteString(", got ")
	builder.WriteString(got)

	return builder.String()
}

// NewErrNotAsExpected creates a new ErrNotAsExpected error.
//
// Parameters:
//   - quote: Whether or not to quote the runes in the error message.
//   - kind: The kind of thing that was not as expected. This is used in the error message.
//   - got: The actual value. If nil, "nothing" is used in the error message.
//   - expecteds: The expected values. If empty, "something" is used in the error message.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"expected <kind> to be <expected>, got <got>"
//
// Where:
//   - <kind>: The kind of thing that was not as expected. This is used in the error message.
//   - <expected>: The expected values. This is used in the error message.
//   - <got>: The actual value. This is used in the error message. If nil, "nothing" is used instead.
func NewErrNotAsExpected(quote bool, kind string, got *rune, expecteds ...rune) error {
	return &ErrNotAsExpected{
		Quote:     quote,
		Kind:      kind,
		Expecteds: expecteds,
		Got:       got,
	}
}

// ErrAfter is an error that occurs after another error.
type ErrAfter struct {
	// Quote is a flag that indicates that the error should be quoted.
	Quote bool

	// Previous is the previous value.
	Previous *rune

	// Inner is the inner error.
	Inner error
}

// Error implements the error interface.
func (e ErrAfter) Error() string {
	var previous string

	if e.Previous == nil {
		previous = "at the start"
	} else if e.Quote {
		previous = strconv.QuoteRune(*e.Previous)
		previous = "after " + previous
	} else {
		previous = string(*e.Previous)
		previous = "after " + previous
	}

	var reason string

	if e.Inner == nil {
		reason = "something went wrong"
	} else {
		reason = e.Inner.Error()
	}

	return previous + ": " + reason
}

// NewErrAfter creates a new ErrAfter error.
//
// Parameters:
//   - quote: A flag indicating whether the previous value should be quoted.
//   - previous: The previous value associated with the error. If not provided, "at the start" is used.
//   - inner: The inner error that occurred. If not provided, "something went wrong" is used.
//
// Returns:
//   - error: The newly created ErrAfter error. Never returns nil.
//
// Format:
//
//	"after <previous>: <inner>"
func NewErrAfter(quote bool, previous *rune, inner error) error {
	return &ErrAfter{
		Quote:    quote,
		Previous: previous,
		Inner:    inner,
	}
}
