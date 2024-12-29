package runes

import (
	"github.com/PlayerR9/mygo-lib/runes/internal"
)

// BytesToUtf8 converts a byte slice into a slice of runes, and returns
// an error if the byte slice contains invalid utf-8 data.
//
// Parameters:
//   - b: The byte slice to convert.
//
// Returns:
//   - []rune: The converted slice of runes.
//   - error: An error if the conversion failed, or nil if it succeeded.
//
// Errors:
//   - ErrInvalidUtf8: If the byte slice contains invalid utf-8 data.
func BytesToUtf8(b []byte) ([]rune, error) {
	if len(b) == 0 {
		return nil, nil
	}

	var chars []rune

	n := internal.BytesToUtf8(b, &chars)
	if n != uint(len(b)) {
		return nil, ErrInvalidUtf8
	}

	return chars, nil
}

// StringToUtf8 converts a string into a slice of runes, and returns
// an error if the string contains invalid utf-8 data.
//
// Parameters:
//   - str: The string to convert.
//
// Returns:
//   - []rune: The converted slice of runes.
//   - error: An error if the conversion failed, or nil if it succeeded.
//
// Errors:
//   - ErrInvalidUtf8: If the string contains invalid utf-8 data.
func StringToUtf8(str string) ([]rune, error) {
	if str == "" {
		return nil, nil
	}

	var chars []rune

	n := internal.StringToUtf8(str, &chars)
	if n != uint(len(str)) {
		return nil, ErrInvalidUtf8
	}

	return chars, nil
}
