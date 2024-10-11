package errors

import "errors"

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
