package common

import (
	"io"
	"os"
)

// TODO: Delete this once go-verify is updated.

// TODO panics with a TODO message. The given message is appended to the
// string "TODO: ". If the message is empty, the message "TODO: Handle this
// case" is used instead.
//
// Parameters:
//   - msg: The message to append to the string "TODO: ".
//
// This function is meant to be used only when the code is being built or
// refactored.
//
// Deprecated: The function was moved to go-verify.
func TODO(msg string) {
	if msg == "" {
		panic("TODO: Handle this case")
	} else {
		panic("TODO: " + msg)
	}
}

// Must is a helper function that wraps a call to a function that returns (T, error) and
// panics if the error is not nil.
//
// This function is intended to be used to handle errors in a way that is easy to read and write.
//
// Parameters:
//   - res: The result of the function.
//   - err: The error returned by the function.
//
// Returns:
//   - T: The result of the function.
//
// Deprecated: The function was moved to go-verify.
func Must[T any](res T, err error) T {
	if err != nil {
		panic(NewErrMust(err))
	}

	return res
}

// WARN prints a warning message to the console.
// The message is prefixed with "[WARNING]:" to indicate its nature.
//
// Parameters:
//   - msg: The warning message to be displayed.
//
// Panics if there is an error writing to the standard output.
//
// Deprecated: The function was moved to go-verify.
func WARN(msg string) {
	data := []byte("[WARNING]: " + msg + "\n")

	n, err := os.Stdout.Write(data)
	if err != nil {
		panic(err)
	} else if n != len(data) {
		panic(io.ErrShortWrite)
	}
}
