package errors

import (
	"errors"
	"strconv"
)

// NewBadParameter creates a new error with the "parameter <param_name> must <must>" message.
//
// Parameters:
//   - param_name: The name of the parameter. If empty, it is omitted.
//   - must: What the parameter must be. If empty, "is invalid" is used instead, otherwise
//     the string is prefixed with "must ".
//
// Returns:
//   - error: The new error. Never returns nil.
func NewBadParameter(param_name, must string) error {
	if param_name != "" {
		param_name = strconv.Quote(param_name) + " "
	}

	if must == "" {
		must = "is invalid"
	} else {
		must = "must " + must
	}

	return errors.New("parameter " + param_name + must)
}
