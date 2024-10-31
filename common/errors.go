package common

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

var (
	// ErrNilReceiver occurs when a method is called on a receiver who was not
	// expected to be nil. This error can be checked with the == operator.
	//
	// Format:
	// 	"receiver must not be nil"
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}

// ErrBadParam occurs when a parameter is bad. (i.e., not a valid value).
type ErrBadParam struct {
	// ParamName is the name of the parameter causing the error.
	ParamName string

	// Msg is the error message describing why the parameter is bad.
	Msg string
}

// Error implements the error interface.
func (e ErrBadParam) Error() string {
	var msg string

	if e.Msg == "" {
		msg = "is invalid"
	} else {
		msg = e.Msg
	}

	if e.ParamName == "" {
		return "parameter " + msg
	} else {
		return "parameter (" + e.ParamName + ") " + msg
	}
}

// NewErrBadParam creates a new ErrBadParam error with the specified parameter name and message.
//
// Parameters:
//   - param_name: The name of the parameter causing the error.
//   - msg: The error message describing why the parameter is bad.
//
// Returns:
//   - error: An instance of ErrBadParam. Never returns nil.
//
// Format:
//
//	"parameter (<param_name>) <msg>"
//
// where:
//   - (<param_name>): The name of the parameter. If empty, it is omitted.
//   - <msg>: The error message describing why the parameter is bad. If empty, "is invalid" is used.
func NewErrBadParam(param_name, msg string) error {
	return &ErrBadParam{
		ParamName: param_name,
		Msg:       msg,
	}
}

// NewErrNilParam is a convenience function that creates a new ErrBadParam error with the specified
// parameter name and the message "must not be nil".
//
// Parameters:
//   - param_name: The name of the parameter causing the error.
//
// Returns:
//   - error: An instance of ErrBadParam. Never returns nil.
//
// Format:
//
//	"parameter (<param_name>) must not be nil"
//
// where:
//   - (<param_name>): The name of the parameter. If empty, it is omitted.
func NewErrNilParam(param_name string) error {
	return &ErrBadParam{
		ParamName: param_name,
		Msg:       "must not be nil",
	}
}

// ErrNotAsExpected occurs when a string is not as expected.
type ErrNotAsExpected struct {
	// Quote if true, the strings will be quoted before being printed.
	Quote bool

	// Kind is the kind of the string that is not as expected.
	Kind string

	// Expecteds are the strings that were expecteds.
	Expecteds []string

	// Got is the actual string.
	Got string
}

// Error implements the error interface.
func (e ErrNotAsExpected) Error() string {
	var kind string

	if e.Kind != "" {
		kind = e.Kind + " to be "
	}

	var got string

	if e.Got == "" {
		got = "nothing"
	} else if e.Quote {
		got = strconv.Quote(e.Got)
	} else {
		got = e.Got
	}

	var builder strings.Builder

	builder.WriteString("expected ")
	builder.WriteString(kind)

	if len(e.Expecteds) > 0 {
		var elems []string

		if !e.Quote {
			elems = e.Expecteds
		} else {
			elems = make([]string, 0, len(e.Expecteds))

			for _, elem := range e.Expecteds {
				str := strconv.Quote(elem)
				elems = append(elems, str)
			}
		}

		builder.WriteString(EitherOrString(elems))
	} else {
		builder.WriteString("something")
	}

	builder.WriteString(", got ")
	builder.WriteString(got)

	return builder.String()
}

// NewErrNotAsExpected creates a new ErrNotAsExpected error.
//
// Parameters:
//   - quote: Whether or not to quote the strings in the error message.
//   - kind: The kind of thing that was not as expected. This is used in the error message.
//   - got: The actual value. If empty, "nothing" is used in the error message.
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
//
// Duplicate values are automatically removed and the list of expected values is sorted in ascending order.
func NewErrNotAsExpected(quote bool, kind string, got string, expecteds ...string) error {
	unique := make([]string, 0, len(expecteds))

	for _, expected := range expecteds {
		pos, ok := slices.BinarySearch(unique, expected)
		if !ok {
			unique = slices.Insert(unique, pos, expected)
		}
	}

	unique = unique[:len(unique):len(unique)]

	return &ErrNotAsExpected{
		Quote:     quote,
		Kind:      kind,
		Expecteds: unique,
		Got:       got,
	}
}

// ErrMust occurs when something must be true.
type ErrMust struct {
	// Inner is the inner error.
	Inner error
}

// Error implements the error interface.
func (e ErrMust) Error() string {
	var msg string

	if e.Inner == nil {
		msg = "something went wrong"
	} else {
		msg = e.Inner.Error()
	}

	return "must: " + msg
}

// NewErrMust creates a new ErrMust error.
//
// Parameters:
//   - inner: The inner error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"must: <inner>"
//
// Where, <inner>: The inner error. If nil, "something went wrong" is used instead.
func NewErrMust(inner error) error {
	return &ErrMust{
		Inner: inner,
	}
}

// Unwrap returns the inner error.
//
// Returns:
//   - error: The inner error.
func (e ErrMust) Unwrap() error {
	return e.Inner
}
