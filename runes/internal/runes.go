package internal

import "unicode/utf8"

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
