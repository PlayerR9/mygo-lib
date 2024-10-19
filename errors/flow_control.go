package errors

import (
	"errors"
)

// try is a utility function that tries to execute a function and returns an error if one occurs.
//
// Parameters:
//   - err: The error pointer. Assumes it is not nil.
//   - fn: The function to execute. Assumes it is not nil.
func try(err *error, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch r := r.(type) {
		case string:
			*err = errors.New(r)
		case error:
			*err = r
		default:
			*err = NewErrPanic(r)
		}
	}()

	fn()
}

// Try executes a panicing function and returns an error if one occurs.
//
// Parameters:
//   - fn: The function to execute.
//
// Returns:
//   - error: The error if one occurs. Otherwise, nil.
func Try(fn func()) error {
	if fn == nil {
		return nil
	}

	var err error

	try(&err, fn)

	return err
}
