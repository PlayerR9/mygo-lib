package runes

import (
	"fmt"
	"slices"
	"unicode/utf8"

	"github.com/PlayerR9/mygo-lib/common"
)

// BytesToUtf8 converts a byte slice to a slice of utf-8 encoded runes. This function stops
// as soon as an error occurs; returning the runes decoded so far.
//
// Parameters:
//   - str: The string to convert.
//
// Returns:
//   - []rune: The slice of runes.
//   - error: An error of type *ErrBadEncoding if the string is not valid UTF-8.
func BytesToUtf8(data []byte) ([]rune, error) {
	if len(data) == 0 {
		return nil, nil
	}

	var chars []rune
	var idx int

	for len(data) > 0 {
		c, size := utf8.DecodeRune(data)
		data = data[size:]

		if c == utf8.RuneError {
			return chars, NewErrBadEncoding(idx)
		}

		idx += size
		chars = append(chars, c)
	}

	return chars, nil
}

// StringToUtf8 converts a string to a slice of utf-8 encoded runes. This function stops
// as soon as an error occurs; returning the runes decoded so far.
//
// Parameters:
//   - str: The string to convert.
//
// Returns:
//   - []rune: The slice of runes.
//   - error: An error of type *ErrBadEncoding if the string is not valid UTF-8.
func StringToUtf8(str string) ([]rune, error) {
	if str == "" {
		return nil, nil
	}

	var chars []rune
	var idx int

	for len(str) > 0 {
		c, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		if c == utf8.RuneError {
			return chars, NewErrBadEncoding(idx)
		}

		idx += size
		chars = append(chars, c)
	}

	return chars, nil
}

// IndicesOf returns a slice of indices that specify where the separator occurs in the data.
//
// Parameters:
//   - data: The data.
//   - sep: The separator.
//
// Returns:
//   - []int: The indices. Nil if the data is empty or the separator is not found.
func IndicesOf(data []rune, sep rune) []int {
	if len(data) == 0 {
		return nil
	}

	var count int

	for i := 0; i < len(data); i++ {
		if data[i] == sep {
			count++
		}
	}

	if count == 0 {
		return nil
	}

	indices := make([]int, 0, count)

	for i := 0; i < len(data); i++ {
		if data[i] == sep {
			indices = append(indices, i)
		}
	}

	return indices
}

// NormalizeRunes is a function that converts '\r\n' to '\n' and tabs to one or more spaces.
//
// Parameters:
//   - chars: The characters to convert.
//   - tab_size: The size of the tab.
//
// Returns:
//   - []rune: The normalized characters.
//   - error: An error if normalization failed.
//
// Errors:
//   - errors.ErrBadParameter: If tab_size is not positive.
//   - errors.ErrAfter: If the characters are not valid UTF-8.
func NormalizeRunes(chars []rune, tab_size int) ([]rune, error) {
	if tab_size <= 0 {
		return chars, common.NewErrBadParameter("tab_size", fmt.Sprintf("must be positive, got %d", tab_size))
	}

	if len(chars) == 0 {
		return nil, nil
	}

	indices := IndicesOf(chars, '\r')

	for _, idx := range indices {
		if idx+1 >= len(chars) {
			return chars, NewErrNotAsExpected(chars[idx], []rune{'\n'}, nil)
		}

		next := chars[idx+1]
		if next != '\n' {
			return chars, NewErrNotAsExpected(chars[idx], []rune{'\n'}, &next)
		}
	}

	for i, idx := range indices {
		chars = slices.Delete(chars, idx-i, idx-i+1)
	}

	var count int

	for i := 0; i < len(chars); i++ {
		if chars[i] == '\t' {
			count++
		}
	}

	if count == 0 {
		return chars, nil
	}

	new_chars := make([]rune, 0, len(chars)+count*(tab_size-1))

	tab := make([]rune, 0, tab_size)
	for i := 0; i < tab_size; i++ {
		tab = append(tab, ' ')
	}

	for _, c := range chars {
		if c == '\t' {
			new_chars = append(new_chars, tab...)
		} else {
			new_chars = append(new_chars, c)
		}
	}

	return new_chars, nil
}

// Repeat is a function that repeats the character.
//
// Parameters:
//   - char: The character to repeat.
//   - count: The number of times to repeat the character.
//
// Returns:
//   - []rune: The repeated character. Returns nil if count is less than or equal to 0.
func Repeat(char rune, count int) []rune {
	if count <= 0 {
		return nil
	}

	chars := make([]rune, 0, count)

	for i := 0; i < count; i++ {
		chars = append(chars, char)
	}

	return chars
}
