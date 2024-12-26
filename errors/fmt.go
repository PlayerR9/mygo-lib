package errors

import "fmt"

// formattedError is an error with a formatted message.
type formattedError struct {
	// format is the format string.
	format string

	// args are the arguments to the format string.
	args []any
}

// Error implements the error interface.
func (fe formattedError) Error() string {
	msg := fmt.Sprintf(fe.format, fe.args...)
	return msg
}

// Errorf returns an error with a formatted message.
//
// Returns:
// 	- error: The new error instance. Never returns nil.
func Errorf(format string, args ...any) error {
	fe := &formattedError{
		format: format,
		args:   args,
	}

	return fe
}

// Unwrap returns the inner errors.
//
// Returns:
// 	- []error: The inner errors.
func (fs formattedError) Unwrap() []error {
	var errs []error

	for _, arg := range fs.args {
		if arg == nil {
			continue
		}

		err, ok := arg.(error)
		if ok {
			errs = append(errs, err)
		}
	}

	return errs
}
