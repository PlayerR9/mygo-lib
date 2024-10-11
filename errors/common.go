package errors

import (
	"errors"
	"strings"
)

var (
	// DefaultErr is the default error that is returned when an error is not provided.
	//
	// Format:
	//   something went wrong
	DefaultErr error
)

func init() {
	DefaultErr = errors.New("something went wrong")
}

// ErrMsgOf is a utility function that returns the error message of the given error.
//
// Parameters:
//   - err: The error.
//
// Returns:
//   - string: The error message of the given error.
//
// Behaviors:
//   - If the given error is nil, the default error message is returned. Thus, if the given error's
//     message is the same as the default message, it will be undistinguishable from a nil error.
func ErrMsgOf(err error) string {
	if err == nil {
		return DefaultErr.Error()
	}

	return err.Error()
}

// EitherOrString is a function that returns a string representation of a slice
// of strings. Empty strings are ignored.
//
// Parameters:
//   - values: The values to convert to a string.
//
// Returns:
//   - string: The string representation.
//
// Example:
//
//	EitherOrString([]string{"a", "b", "c"}) // "either a, b, or c"
func EitherOrString(elems []string) string {
	var str string

	switch len(elems) {
	case 0:
		// Do nothing
	case 1:
		str = elems[0]
	case 2:
		str = "either " + elems[0] + " or " + elems[1]
	default:
		str = "either " + strings.Join(elems[:len(elems)-1], ", ") + ", or " + elems[len(elems)-1]
	}

	return str
}
