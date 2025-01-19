package errors

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PlayerR9/mygo-lib/errors/internal"
)

var (
	// ErrNilReceiver occurs when a method is called on a receiver that was not expected to be
	// nil. This error can be checked with the == operator.
	//
	// Format:
	//
	//		"receiver must not be nil"
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}

// ErrBadParam occurs when a function is called with a parameter that does not
// satisfy the function's contract.
type ErrBadParam struct {
	// ParamName is the name of the parameter that caused the error.
	ParamName string

	// Message is the error message.
	Message string
}

// Error implements error.
func (e ErrBadParam) Error() string {
	var msg string

	if e.Message == "" {
		msg = "is not valid"
	} else {
		msg = e.Message
	}

	if e.ParamName == "" {
		return "parameter " + msg
	} else {
		return "parameter (" + e.ParamName + ") " + msg
	}
}

// NewErrBadParam creates a new ErrBadParam error with the given parameter name and
// message.
//
// Parameters:
//   - param_name: The name of the parameter that caused the error.
//   - msg: The error message.
//
// Returns:
//   - error: The new ErrBadParam error. Never returns nil.
//
// Format:
//
//	"parameter (<param_name>) <msg>"
//
// Where:
//   - <param_name> is the name of the parameter that caused the error.
//   - <msg> is the error message. If empty, the message defaults to "is not valid".
//
// However, if param_name is empty, the format is:
//
//	"parameter <msg>"
//
// Where, <msg> is the error message. If empty, the message defaults to "is not valid".
func NewErrBadParam(param_name, msg string) error {
	err := &ErrBadParam{
		ParamName: param_name,
		Message:   msg,
	}

	return err
}

// NewErrNilParam is a convenience function for creating a new ErrBadParam error
// with the message "must not be nil".
//
// Parameters:
//   - param_name: The name of the parameter that caused the error.
//
// Returns:
//   - error: The new ErrBadParam error. Never returns nil.
//
// Format:
//
//	"parameter (<param_name>) must not be nil"
//
// Where, <param_name> is the name of the parameter that caused the error.
//
// However, if param_name is empty, the format is:
//
//	"parameter must not be nil"
func NewErrNilParam(param_name string) error {
	err := &ErrBadParam{
		ParamName: param_name,
		Message:   "must not be nil",
	}

	return err
}

// ErrUnexpected occurs when an unexpected value is encountered.
type ErrUnexpected struct {
	// Kind is the kind of unexpected value.
	Kind string

	// Got is the unexpected value.
	Got string

	// Want is the expected value.
	Want string
}

// Error implements error.
func (eu ErrUnexpected) Error() string {
	var got string

	if eu.Got == "" {
		got = "nothing"
	} else {
		got = eu.Got
	}

	var want string

	if eu.Want == "" {
		want = "something"
	} else {
		want = eu.Want
	}

	if eu.Kind == "" {
		return "want " + want + ", got " + got
	} else {
		var builder strings.Builder

		_, _ = builder.WriteString("want ")
		_, _ = builder.WriteString(eu.Kind)
		_, _ = builder.WriteString(" to be ")
		_, _ = builder.WriteString(want)
		_, _ = builder.WriteString(", got ")
		_, _ = builder.WriteString(got)

		return builder.String()
	}
}

// NewErrUnexpected creates a new ErrUnexpected error with the given kind, want and got values.
//
// Parameters:
//   - kind: The kind of unexpected value.
//   - want: The expected value.
//   - got: The unexpected value.
//
// Returns:
//   - error: The new ErrUnexpected error. Never returns nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of unexpected value.
//   - <want> is the expected value. If empty, the value defaults to "something".
//   - <got> is the unexpected value. If empty, the value defaults to "nothing".
//
// However, if kind is empty, the format is:
//
//	"want <want> to be <got>"
//
// Where:
//   - <want> is the expected value. If empty, the value defaults to "something".
//   - <got> is the unexpected value. If empty, the value defaults to "nothing".
func NewErrUnexpected(kind, want, got string) error {
	err := &ErrUnexpected{
		Kind: kind,
		Got:  got,
		Want: want,
	}

	return err
}

// NewErrUnexpectedQuoted creates a new ErrUnexpected error with the given kind, want and got values,
// after quoting the got value and the want values.
//
// Parameters:
//   - kind: The kind of unexpected value.
//   - got: The unexpected value.
//   - wants: The expected values.
//
// Returns:
//   - error: The new ErrUnexpected error. Never returns nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of unexpected value.
//   - <want> is the expected value. If empty, the value defaults to "something".
//   - <got> is the unexpected value. If empty, the value defaults to "nothing".
//
// However, if kind is empty, the format is:
//
//	"want <want> to be <got>"
//
// Where:
//   - <want> is the expected value. If empty, the value defaults to "something".
//   - <got> is the unexpected value. If empty, the value defaults to "nothing".
func NewErrUnexpectedQuoted(kind string, got string, wants ...string) error {
	var want string

	if len(wants) == 0 {
		want = ""
	} else {
		wants = internal.RejectEmpty(wants)
		internal.Quote(wants)

		want = internal.EitherOr(wants)
	}

	if got != "" {
		got = strconv.Quote(got)
	}

	err := &ErrUnexpected{
		Kind: kind,
		Got:  got,
		Want: want,
	}

	return err
}
