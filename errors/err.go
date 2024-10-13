package errors

import (
	"strconv"

	faults "github.com/PlayerR9/mygo-lib/PlayerR9/go-fault"
)

// ErrorCode is an error code.
type ErrorCode int

const (
	// BadParameter is an error code used to indicate that a parameter is invalid.
	BadParameter ErrorCode = iota
)

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
		return faults.New(BadParameter, "parameter ("+strconv.Quote(param)+") "+msg)
	} else {
		return faults.New(BadParameter, "parameter "+msg)
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
		return faults.New(BadParameter, "parameter ("+strconv.Quote(param)+") must not be nil")
	} else {
		return faults.New(BadParameter, "parameter must not be nil")
	}
}
