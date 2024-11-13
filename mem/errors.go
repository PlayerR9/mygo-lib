package mem

import (
	"fmt"
	"strconv"
)

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

// ErrInvalidType occurs when a type is not as expected.
type ErrInvalidType struct {
	// Type is the expected type.
	Type any

	// Got is the actual type.
	Got any
}

// Error implements the error interface.
func (e ErrInvalidType) Error() string {
	var expected string

	expected = fmt.Sprintf("%T", e.Type)

	var got string

	if e.Got == nil {
		got = "<nil>"
	} else {
		got = fmt.Sprintf("%T", e.Got)
	}

	return "want " + expected + ", got " + got
}

// NewErrInvalidType creates a new ErrInvalidType error with the specified expected type and actual type.
//
// Parameters:
//   - want: The expected type.
//   - got: The actual type.
//
// Returns:
//   - error: The new ErrInvalidType error. Never returns nil.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want>: The expected type.
//   - <got>: The actual type. If nil, "<nil>" is used instead.
func NewErrInvalidType(got any, want any) error {
	return &ErrInvalidType{
		Type: want,
		Got:  got,
	}
}
