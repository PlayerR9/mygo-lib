package common

import (
	"errors"
	"fmt"
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

// ErrNotAsExpected occurs when a value is not as expected.
type ErrNotAsExpected[T any] struct {
	// Quote if true, the values will be quoted before being printed.
	Quote bool

	// Kind is the kind of the value that is not as expected.
	Kind string

	// Expected are the values that were expected.
	Expecteds []T

	// Got is the actual value.
	Got any
}

// Error implements the error interface.
func (e ErrNotAsExpected[T]) Error() string {
	var kind string

	if e.Kind != "" {
		kind = e.Kind + " to be "
	}

	var got string

	if e.Got == nil {
		got = "nothing"
	} else if e.Quote {
		got = fmt.Sprint(e.Got)
		got = strconv.Quote(got)
	} else {
		got = fmt.Sprint(e.Got)
	}

	var builder strings.Builder

	builder.WriteString("expected ")
	builder.WriteString(kind)

	elems := make([]string, 0, len(e.Expecteds))

	for _, elem := range e.Expecteds {
		str := fmt.Sprint(elem)
		elems = append(elems, str)
	}

	if e.Quote {
		for i := range elems {
			elems[i] = strconv.Quote(elems[i])
		}
	}

	builder.WriteString(EitherOrString(elems))
	builder.WriteString(", got ")
	builder.WriteString(got)

	return builder.String()
}

// NewErrNotAsExpected creates a new ErrNotAsExpected error.
//
// Parameters:
//   - quote: Whether or not to quote the values in the error message.
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
func NewErrNotAsExpected[T any](quote bool, kind string, got any, expecteds ...T) error {
	return &ErrNotAsExpected[T]{
		Quote:     quote,
		Kind:      kind,
		Expecteds: expecteds,
		Got:       got,
	}
}
