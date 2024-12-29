package internal

import (
	"io"
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
//   - io.ErrShortWrite: If the write operation writes fewer bytes than expected.
//   - any other error: Implementation-specific error.
func WriteBytes(w io.Writer, data []byte) error {
	n, err := w.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}
