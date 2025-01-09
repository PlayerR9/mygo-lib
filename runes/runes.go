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

// Split divides a slice of runes into sub-slices based on a separator rune.
//
// Parameters:
//   - chars: The slice of runes to be split.
//   - sep: The rune used as the separator.
//
// Returns:
//   - [][]rune: A slice of rune slices, where each sub-slice contains runes between separators.
func Split(chars []rune, sep rune) [][]rune {
	if len(chars) == 0 {
		return nil
	}

	result := internal.Split(chars, sep)
	return result
}

// TrimLeft trims whitespace from the left of a slice of runes.
//
// Parameters:
//   - chars: The slice of runes to trim.
//
// Returns:
//   - uint: The number of characters trimmed from the left.
func TrimLeft(chars *[]rune, fn func(rune) bool) uint {
	if fn == nil || chars == nil || len(*chars) == 0 {
		return 0
	}

	n := internal.TrimLeft(chars, fn)
	return n
}

// TrimRight trims whitespace from the right of a slice of runes.
//
// Parameters:
//   - chars: The slice of runes to trim.
//
// Returns:
//   - uint: The number of characters trimmed from the right.
func TrimRight(chars *[]rune, fn func(rune) bool) uint {
	if fn == nil || chars == nil || len(*chars) == 0 {
		return 0
	}

	n := internal.TrimRight(chars, fn)
	return n
}

// TrimSpace trims whitespace from the beginning and end of a slice of runes.
//
// Parameters:
//   - chars: The slice of runes to trim.
//
// Returns:
//   - uint: The number of characters trimmed from the beginning and end.
func TrimSpace(chars *[]rune) uint {
	if chars == nil || len(*chars) == 0 {
		return 0
	}

	n := internal.TrimSpace(chars)
	return n
}

// Equal checks if two slices of runes are equal.
//
// Parameters:
//   - a: The first slice of runes.
//   - b: The second slice of runes.
//
// Returns:
//   - bool: True if the slices are equal, false otherwise.
func Equal(a, b []rune) bool {
	ok := internal.Equal(a, b)
	return ok
}
