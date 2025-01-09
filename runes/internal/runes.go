package internal

import (
	"unicode"
	"unicode/utf8"
)

// BytesToUtf8 converts a byte slice into a slice of runes, and returns
// an error if the byte slice contains invalid utf-8 data.
//
// Parameters:
//   - b: The byte slice to convert.
//   - chars: The slice to append the converted runes to.
//
// Returns:
//   - uint: The total number of bytes consumed.
func BytesToUtf8(b []byte, chars *[]rune) uint {
	var total uint

	for len(b) > 0 {
		r, s := utf8.DecodeRune(b)
		if r == utf8.RuneError {
			break
		}

		*chars = append(*chars, r)
		b = b[s:]
		total += uint(s)
	}

	return total
}

// StringToUtf8 converts a string into a slice of runes, and returns
// an error if the string contains invalid utf-8 data.
//
// Parameters:
//   - str: The string to convert.
//   - chars: The slice to append the converted runes to.
//
// Returns:
//   - uint: The total number of bytes consumed.
func StringToUtf8(str string, chars *[]rune) uint {
	var total uint

	for len(str) > 0 {
		r, s := utf8.DecodeRuneInString(str)
		if r == utf8.RuneError {
			break
		}

		*chars = append(*chars, r)
		str = str[s:]
		total += uint(s)
	}

	return total
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
	var result [][]rune
	var current []rune

	for _, char := range chars {
		if char == sep {
			if len(current) > 0 {
				result = append(result, current)
				current = nil
			}
		} else {
			current = append(current, char)
		}
	}

	if len(current) > 0 {
		result = append(result, current)
	}

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
	lenChars := uint(len(*chars))

	var i uint

	for i < lenChars && fn((*chars)[i]) {
		i++
	}

	if i == lenChars {
		clear(*chars)
		*chars = nil
	} else if i > 0 {
		clear((*chars)[:i])
		*chars = (*chars)[i:]
	}

	return i
}

// TrimRight trims whitespace from the right of a slice of runes.
//
// Parameters:
//   - chars: The slice of runes to trim.
//
// Returns:
//   - uint: The number of characters trimmed from the right.
func TrimRight(chars *[]rune, fn func(rune) bool) uint {
	lenChars := uint(len(*chars))

	i := lenChars

	for i > 0 && fn((*chars)[i-1]) {
		i--
	}

	if i == 0 {
		clear(*chars)
		*chars = nil
	} else if i < lenChars {
		clear((*chars)[i:])
		*chars = (*chars)[:i]
	}

	return lenChars - i
}

// TrimSpace trims whitespace from the beginning and end of a slice of runes.
//
// Parameters:
//   - chars: The slice of runes to trim.
//
// Returns:
//   - uint: The number of characters trimmed from the beginning and end.
func TrimSpace(chars *[]rune) uint {
	var count uint

	n := TrimLeft(chars, unicode.IsSpace)
	count += n

	if len(*chars) == 0 {
		return count
	}

	n = TrimRight(chars, unicode.IsSpace)
	count += n

	return count
}

// Equal checks if two slices of runes are equal.
//
// Parameters:
//   - a: The first slice to compare.
//   - b: The second slice to compare.
//
// Returns:
//   - bool: True if the slices are equal, false otherwise.
func Equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i, c := range a {
		if c != b[i] {
			return false
		}
	}

	return true
}
