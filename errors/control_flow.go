package errors

// Yeet panics if the provided error is not nil.
//
// Parameters:
//   - err: The error to check. If it is not nil, the function will panic.
//
// This is a joke function name for the "throw" keyword.
func Yeet(err error) {
	if err == nil {
		return
	}

	panic(err)
}
