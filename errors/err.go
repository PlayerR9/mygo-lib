package errors

import (
	"strconv"
)

// ErrorCode is an error code.
type ErrorCode int

const (
	// BadParameter is an error code used to indicate that a parameter is invalid.
	BadParameter ErrorCode = iota
)

// Err is a generic error.
type Err struct {
	// Code is the code of the error.
	Code ErrorCode

	// Msg is the message of the error.
	Msg string
}

// Error implements the error interface.
func (e Err) Error() string {
	if e.Msg == "" {
		return "(" + e.Code.String() + ") something went wrong"
	} else {
		return "(" + e.Code.String() + ") " + e.Msg
	}
}

// NewErr creates a new error.
//
// Parameters:
//   - code: The code of the error.
//   - msg: The message of the error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	(<code>) <msg>
//
// Where:
//   - <code>: The code of the error.
//   - <msg>: The message of the error. If empty, "something went wrong" is used instead.
func NewErr(code ErrorCode, msg string) error {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}

// NewErrBadParameter is a convenience function that creates a new Err with BadParameter code and msg.
//
// Parameters:
//   - param: The name of the parameter.
//   - msg: What the parameter must be. If empty, "is invalid" is used instead.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// The message must be written assuming the "must " precedes it.
func NewErrBadParameter(param string, msg string) error {
	if msg != "" {
		msg = "must " + msg
	} else {
		msg = "is invalid"
	}

	if param != "" {
		return &Err{
			Code: BadParameter,
			Msg:  "parameter (" + strconv.Quote(param) + ") " + msg,
		}
	} else {
		return &Err{
			Code: BadParameter,
			Msg:  "parameter " + msg,
		}
	}
}

// NewErrNilParameter is a convenience function that is like NewErrBadParameter but for nil parameters.
//
// Parameters:
//   - param: The name of the parameter.
//
// Returns:
//   - error: The new error. Never returns nil.
func NewErrNilParameter(param string) error {
	if param != "" {
		return &Err{
			Code: BadParameter,
			Msg:  "parameter (" + strconv.Quote(param) + ") must not be nil",
		}
	} else {
		return &Err{
			Code: BadParameter,
			Msg:  "parameter must not be nil",
		}
	}
}
