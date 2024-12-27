package io

import (
	"io"

	"github.com/PlayerR9/mygo-lib/errors"
)

// WriteString writes the given string to the writer and returns an error if the write operation fails.
//
// Parameters:
//   - w: The writer to write the string to. Must not be nil.
//   - str: The string to write.
//
// Returns:
//   - error: An error if the write operation fails.
//
// Errors:
//   - errors.ErrNilParam: If the parameter is nil.
//   - io.ErrShortWrite: If the write operation fails.
func WriteString(w io.Writer, str string) error {
	if w == nil {
		return errors.NewErrNilParam("w")
	} else if str == "" {
		return nil
	}

	data := []byte(str)

	n, err := w.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}
