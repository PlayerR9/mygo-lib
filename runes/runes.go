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
	lenB := uint(len(b))
	if lenB == 0 {
		return nil, nil
	}

	var chars []rune

	n := internal.BytesToUtf8(b, &chars)
	if n != lenB {
		return chars, ErrInvalidUtf8
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

// Split divides a slice of runes into sub-slices based on a separator function.
//
// Parameters:
//   - chars: The slice of runes to be split.
//   - isSep: The function used to determine if a rune is a separator.
//
// Returns:
//   - [][]rune: A slice of rune slices, where each sub-slice contains runes between separators.
//
// Panics:
//   - ErrNoPredicate: If the separator function is nil.
func Split(chars []rune, isSep func(rune) bool) [][]rune {
	if isSep == nil {
		panic(ErrNoPredicate)
	}

	if len(chars) == 0 {
		return nil
	}

	result := internal.Split(chars, isSep)
	return result
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
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	for i, c := range a {
		if c != b[i] {
			return false
		}
	}

	return true
}
