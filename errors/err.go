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

// NewBadParameter is a convenience function that creates a new Fault with BadParameter code and msg.
//
// Parameters:
//   - param_name: The name of the parameter. If empty, it is omitted.
//   - must: What the parameter must be. If empty, "is invalid" is used instead, otherwise
//     the string is prefixed with "must ".
//
// Returns:
//   - Fault: The new fault. Never returns nil.
func NewBadParameter(param_name, must string) faults.Fault {
	if param_name != "" {
		param_name = strconv.Quote(param_name) + " "
	}

	if must == "" {
		must = "is invalid"
	} else {
		must = "must " + must
	}

	return faults.New(BadParameter, "Parameter "+param_name+must)
}
