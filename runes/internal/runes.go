package internal

import (
	"unicode/utf8"
)

// BytesToUtf8 converts a byte slice into a slice of runes, and returns
// an error if the byte slice contains invalid utf-8 data.
//
// Parameters:
//   - data: The byte slice to convert.
//   - dest: The slice to append the converted runes to.
//
// Returns:
//   - uint: The total number of bytes consumed.
func BytesToUtf8(data []byte, dest *[]rune) uint {
	var total uint

	for len(data) > 0 {
		r, s := utf8.DecodeRune(data)
		if r == utf8.RuneError {
			break
		}

		data = data[s:]
		total += uint(s)

		*dest = append(*dest, r)
	}

	return total
}

// StringToUtf8 converts a string into a slice of runes, and returns
// an error if the string contains invalid utf-8 data.
//
// Parameters:
//   - str: The string to convert.
//   - dest: The slice to append the converted runes to.
//
// Returns:
//   - uint: The total number of bytes consumed.
func StringToUtf8(str string, dest *[]rune) uint {
	var total uint

	for len(str) > 0 {
		r, s := utf8.DecodeRuneInString(str)
		if r == utf8.RuneError {
			break
		}

		str = str[s:]
		total += uint(s)

		*dest = append(*dest, r)
	}

	return total
}

// Split divides a slice of runes into sub-slices based on a separator rune.
//
// Parameters:
//   - s: The slice of runes to be split.
//   - isSep: The rune used as the separator.
//
// Returns:
//   - [][]rune: A slice of rune slices, where each sub-slice contains runes between separators.
func Split(s []rune, isSep func(rune) bool) [][]rune {
	var result [][]rune
	var current []rune

	for _, char := range s {
		ok := isSep(char)
		if !ok {
			current = append(current, char)
		} else if len(current) > 0 {
			result = append(result, current)
			current = nil
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
//   - s: The slice of runes to trim.
//   - isTrim: The function used to determine if a rune should be trimmed.
//
// Returns:
//   - uint: The number of characters trimmed from the left.
func TrimLeft(s *[]rune, isTrim func(rune) bool) uint {
	lenChars := uint(len(*s))

	var end uint

	for end < lenChars {
		ok := isTrim((*s)[end])
		if !ok {
			break
		}

		end++
	}

	if end == lenChars {
		clear(*s)
		*s = nil
	} else if end > 0 {
		clear((*s)[:end])
		*s = (*s)[end:]
	}

	return end
}

// TrimRight trims whitespace from the right of a slice of runes.
//
// Parameters:
//   - s: The slice of runes to trim.
//   - isTrim: The function used to determine if a rune should be trimmed.
//
// Returns:
//   - uint: The number of characters trimmed from the right.
func TrimRight(s *[]rune, isTrim func(rune) bool) uint {
	lenChars := uint(len(*s))

	end := lenChars

	for end > 0 {
		ok := isTrim((*s)[end-1])
		if !ok {
			break
		}

		end--
	}

	if end == 0 {
		clear(*s)
		*s = nil
	} else if end < lenChars {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return lenChars - end
}

// Trim trims characters from the beginning and end of a slice of runes.
//
// Parameters:
//   - s: The slice of runes to trim.
//   - isTrim: The function used to determine if a rune should be trimmed.
//
// Returns:
//   - uint: The number of characters trimmed from the beginning and end.
func Trim(s *[]rune, isTrim func(rune) bool) uint {
	var count uint

	n := TrimLeft(s, isTrim)
	count += n

	if len(*s) == 0 {
		return count
	}

	n = TrimRight(s, isTrim)
	count += n

	return count
}
