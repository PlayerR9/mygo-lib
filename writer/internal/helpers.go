package internal

// WriteBytes writes the given byte slice to the provided writer.
//
// Parameters:
//   - w: The writer to which data will be written. Must not be nil.
//   - data: The bytes to write.
//
// Returns:
//   - error: An error if the write operation fails.
//
// Errors:
//   - any error: Implementation-specific error.
func WriteBytes(w interface {
	Write(p []byte) (int, error)
}, data []byte) error {
	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		panic(BugMismatchWritten)
	}

	return nil
}
