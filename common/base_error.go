package common

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

// New returns a new error with the given message. If the message is empty, it
// defaults to DefaultErrorMessage.
//
// Parameters:
//   - msg: The error message.
//
// Returns:
//   - error: The new error. Never returns nil.
func New(msg string) error {
	if msg == "" {
		msg = DefaultErrorMessage
	}

	be := &baseError{
		msg: msg,
	}

	return be
}
