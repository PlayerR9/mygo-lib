package io

import (
	"io"

	common "github.com/PlayerR9/mygo-lib/common"
)

// WriteBytes writes the given byte slice to the provided writer.
//
// Parameters:
//   - w: The writer to which data will be written. Must not be nil.
//   - data: The bytes to write. If empty, the function returns nil.
//
// Returns:
//   - error: An error if the write operation fails.
//
// Errors:
//   - common.NewErrNilParam: If the writer is nil.
//   - io.ErrShortWrite: If the write operation writes fewer bytes than expected.
//   - any other error: Implementation-specific error.
func WriteBytes(w io.Writer, data []byte) error {
	if w == nil {
		return common.NewErrNilParam("w")
	} else if len(data) == 0 {
		return nil
	}

	n, err := w.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

// WriteString writes the given string to the writer and returns an error if the write operation fails.
//
// Parameters:
//   - w: The writer to write the string to. Must not be nil.
//   - str: The string to write. If empty, the function returns nil.
//
// Returns:
//   - error: An error if the write operation fails.
//
// Errors:
//   - common.NewErrNilParam: If the writer is nil.
//   - io.ErrShortWrite: If the write operation writes fewer bytes than expected.
//   - any other error: Implementation-specific error.
func WriteString(w io.Writer, str string) error {
	if w == nil {
		return common.NewErrNilParam("w")
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
