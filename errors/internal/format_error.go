package internal

import (
	"fmt"
)

// formatError is an error with a formatted message.
type formatError struct {
	// format is the format string.
	format string

	// args are the arguments to the format string.
	args []any
}

// Error implements the error interface.
func (fe formatError) Error() string {
	msg := fmt.Sprintf(fe.format, fe.args...)
	return msg
}

// NewFormatError returns an error with a formatted message.
//
// Returns:
//   - error: The new error instance. Never returns nil.
func NewFormatError(format string, args []any) error {
	fe := &formatError{
		format: format,
		args:   args,
	}

	return fe
}

// Unwrap returns the inner errors.
//
// Returns:
//   - []error: The inner errors.
func (fs formatError) Unwrap() []error {
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
