package runes

import (
	"unicode/utf8"
)

// BytesToUtf8 converts a byte slice to a slice of runes.
//
// Parameters:
//   - data: The byte slice to convert.
//
// Returns:
//   - []rune: The slice of runes.
//   - error: An error if conversion failed.
//
// Errors:
//   - ErrBadEncoding: If the byte slice is not valid utf-8.
func BytesToUtf8(data []byte) ([]rune, error) {
	if len(data) == 0 {
		return nil, nil
	}

	var chars []rune

	for len(data) > 0 {
		c, size := utf8.DecodeRune(data)
		if c == utf8.RuneError {
			return nil, ErrBadEncoding
		}

		data = data[size:]
		chars = append(chars, c)
	}

	chars = chars[:len(chars):len(chars)]

	return chars, nil
}

// StringToUtf8 is like BytesToUtf8 but for strings.
//
// Parameters:
//   - str: The string to convert.
//
// Returns:
//   - []rune: The slice of runes.
//   - error: An error if conversion failed.
//
// Errors:
//   - ErrBadEncoding: If the string is not valid utf-8.
func StringToUtf8(str string) ([]rune, error) {
	if str == "" {
		return nil, nil
	}

	var chars []rune

	for len(str) > 0 {
		c, size := utf8.DecodeRuneInString(str)
		if c == utf8.RuneError {
			return nil, ErrBadEncoding
		}

		str = str[size:]
		chars = append(chars, c)
	}

	chars = chars[:len(chars):len(chars)]

	return chars, nil
}

// Repeat generates a slice of runes by repeating a given character a specified number of times.
//
// Parameters:
//   - char: The character to repeat.
//   - count: The number of times to repeat the character.
//
// Returns:
//   - []rune: The resulting slice of repeated characters. Nil if count is less than or equal to 0.
func Repeat(char rune, count int) []rune {
	if count <= 0 {
		return nil
	}

	slice := make([]rune, 0, count)

	for i := 0; i < count; i++ {
		slice = append(slice, char)
	}

	return slice
}
