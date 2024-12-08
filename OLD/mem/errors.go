package mem

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

// ErrInvalidObject occurs when an object is no longer valid, especially
// when the method `Free()` is called.
type ErrInvalidObject struct {
	// MethodName is the name of the method that caused the error.
	MethodName string
}

// Error implements the error interface.
func (e ErrInvalidObject) Error() string {
	return "cannot call method " + strconv.Quote(e.MethodName) + ": object is no longer valid"
}

// NewErrInvalidObject creates a new ErrInvalidObject error.
//
// Parameters:
//   - method_name: The name of the method that caused the error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"cannot call method <method_name>: object is no longer valid"
func NewErrInvalidObject(method_name string) error {
	return &ErrInvalidObject{
		MethodName: method_name,
	}
}

// ErrRelease occurs when the `Release()` function fails.
type ErrRelease struct {
	// Target is the target of the function.
	Targets []string

	// Inner is the reason why the function failed.
	Inner error
}

// Error implements error.
func (e ErrRelease) Error() string {
	var msg string

	if e.Inner == nil {
		msg = "something went wrong"
	} else {
		msg = e.Inner.Error()
	}

	if len(e.Targets) == 0 {
		return "Release() = " + msg
	}

	targets := make([]string, len(e.Targets))
	copy(targets, e.Targets)

	slices.Reverse(targets)

	return "Release(" + strings.Join(targets, " -> ") + ") = " + msg
}

// NewErrRelease creates a new ErrRelease error.
//
// Parameters:
//   - target: The innermost target.
//   - inner: The reason why the function failed.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"Release(<target>) = <inner>"
//
// Where:
//   - <target>: The target of the function.
//   - <inner>: The error returned by the `Release()` method. If nil, "something went wrong" is used instead.
func NewErrRelease(target string, inner error) error {
	return &ErrRelease{
		Targets: []string{target},
		Inner:   inner,
	}
}

// Unwrap returns the inner error.
//
// Returns:
//   - error: The inner error.
func (e ErrRelease) Unwrap() error {
	return e.Inner
}

// AppendTarget appends the given target to the list of targets.
//
// Returns:
//   - error: ErrNilReceiver if the receiver is nil.
func (e *ErrRelease) AppendTarget(target string) error {
	if e == nil {
		return ErrNilReceiver
	}

	e.Targets = append(e.Targets, target)

	return nil
}
