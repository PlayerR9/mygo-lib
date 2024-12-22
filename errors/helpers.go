package errors

// ErrorMessageOf returns the error message of the given error, or "something went wrong" if the error is nil.
//
// Parameters:
//   - err: The error to get the message of.
//
// Returns:
//   - string: The error message of the given error, or "something went wrong" if the error is nil.
func ErrorMessageOf(err error) string {
	if err == nil {
		return "something went wrong"
	}

	msg := err.Error()
	return msg
}
