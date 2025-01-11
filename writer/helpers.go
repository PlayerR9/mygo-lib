package io

import (
	"unicode/utf8"

	"github.com/PlayerR9/mygo-lib/writer/internal"
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
//   - ErrNoWriter: If the writer is nil.
//   - any other error: Implementation-specific error.
func WriteBytes(w Writer, data []byte) error {
	if w == nil {
		return ErrNoWriter
	} else if len(data) == 0 {
		return nil
	}

	err := internal.WriteBytes(w, data)
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
//   - ErrNoWriter: If the writer is nil.
//   - any other error: Implementation-specific error.
func WriteString(w Writer, str string) error {
	if w == nil {
		return ErrNoWriter
	} else if str == "" {
		return nil
	}

	err := internal.WriteBytes(w, []byte(str))
	return err
}

// WriteRune writes the given rune to the writer and returns an error if the write operation fails.
//
// Parameters:
//   - w: The writer to write the rune to. Must not be nil.
//   - r: The rune to write.
//
// Returns:
//   - error: An error if the write operation fails.
//
// Errors:
//   - ErrNoWriter: If the writer is nil.
//   - any other error: Implementation-specific error.
func WriteRune(w Writer, r rune) error {
	if w == nil {
		return ErrNoWriter
	}

	data := utf8.AppendRune(nil, r)

	err := internal.WriteBytes(w, data)
	return err
}
