package runes

import (
	"slices"
	"unicode/utf8"

	"github.com/PlayerR9/mygo-lib/common"
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
//   - []rune: The resulting slice of repeated characters. Nil if count is zero.
func Repeat(char rune, count uint) []rune {
	if count == 0 {
		return nil
	}

	slice := make([]rune, 0, count)

	for i := uint(0); i < count; i++ {
		slice = append(slice, char)
	}

	return slice
}

// NormalizeNewlines takes a slice of runes and normalizes any instances of "\r\n" into "\n".
//
// Parameters:
//   - chars: The characters to normalize.
//
// Returns:
//   - error: An error if normalization failed.
//
// Errors:
//   - ErrAt: If '\r' is not followed by '\n' at the specified index. This error wraps
//     ErrNotAsExpected.
func normalizeNewlines(chars *[]rune) error {
	indices := IndicesOf(*chars, '\r')
	if len(indices) == 0 {
		return nil
	}

	next_idx := indices[len(indices)-1] + 1

	if next_idx >= uint(len(*chars)) {
		return common.NewErrAt(next_idx, NewErrNotAsExpected(true, "", nil, '\n'))
	} else if (*chars)[next_idx] != '\n' {
		return common.NewErrAt(next_idx, NewErrNotAsExpected(true, "", &(*chars)[next_idx], '\n'))
	}

	*chars = slices.Delete(*chars, int(next_idx-1), int(next_idx))

	var offset uint

	for _, idx := range indices[:len(indices)-1] {
		idx += 1 - offset
		char := (*chars)[idx]

		if char != '\n' {
			return common.NewErrAt(idx, NewErrNotAsExpected(true, "", &char, '\n'))
		}

		*chars = slices.Delete(*chars, int(idx-1), int(idx))
		offset++
	}

	return nil
}

// normalizeTabs replaces all tabs in chars with repl.
//
// The function has no side effects other than modifying chars.
func normalizeTabs(chars *[]rune, repl []rune) {
	indices := IndicesOf(*chars, '\t')
	if len(indices) == 0 {
		return
	}

	offset := uint(len(repl) - 1)
	var delta uint

	for _, idx := range indices {
		idx += delta

		*chars = slices.Delete(*chars, int(idx), int(idx+1))
		*chars = slices.Insert(*chars, int(idx), repl...)

		delta += offset
	}
}

// Normalize normalizes the runes in chars by replacing all "\r\n" with "\n" and
// all "\t" with the appropriate number of spaces depending on tab_size.
//
// The function normalizes the runes in place and has no other side effects.
//
// Parameters:
//   - chars: The characters to normalize.
//   - tab_size: The size of the tab stop.
//
// Returns:
//   - error: An error if normalization fails.
//
// Errors:
//   - common.ErrBadParam: If tab_size is not positive.
//   - ErrAt: If '\r' is not followed by '\n' at the specified index. This error wraps
//     ErrNotAsExpected.
func Normalize(chars *[]rune, tab_size uint) error {
	if chars == nil || len(*chars) == 0 {
		return nil
	} else if tab_size == 0 {
		return common.NewErrBadParam("tab_size", "must not be zero")
	}

	err := normalizeNewlines(chars)
	if err != nil {
		return err
	}

	repl := Repeat(' ', tab_size)

	normalizeTabs(chars, repl)

	return nil
}

// Equals checks if two slices of runes are equal.
//
// The two slices are considered equal if they have the same length and the same
// elements in the same order.
//
// Parameters:
//   - first: The first slice of runes.
//   - second: The second slice of runes.
//
// Returns:
//   - bool: True if the two slices are equal, false otherwise.
func Equals(first, second []rune) bool {
	if len(first) != len(second) {
		return false
	} else if len(first) == 0 {
		return true
	}

	for i, c := range first {
		if c != second[i] {
			return false
		}
	}

	return true
}
