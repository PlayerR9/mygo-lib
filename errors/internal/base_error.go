package internal

// THIS IS DONE TO AVOID DEPENDENCY TO 'errors' PACKAGE.

// baseError is the base error type.
type baseError struct {
	// msg is the error message.
	msg string
}

// Error implements error.
func (be baseError) Error() string {
	return be.msg
}

// NewBaseError returns a new error with the given message.
//
// Parameters:
//   - msg: The error message.
//
// Returns:
//   - error: The new error. Never returns nil.
func NewBaseError(msg string) error {
	be := &baseError{
		msg: msg,
	}

	return be
}
