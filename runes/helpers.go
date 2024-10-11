package runes

import (
	"strconv"
	"strings"
)

// QuoteRunes returns a slice of quoted strings from a slice of runes.
//
// Parameters:
//   - slice: The slice of runes.
//
// Returns:
//   - []string: The slice of quoted strings. Nil if the slice is empty.
func QuoteRunes(slice []rune) []string {
	if len(slice) == 0 {
		return nil
	}

	elems := make([]string, 0, len(slice))

	for _, elem := range slice {
		elems = append(elems, strconv.QuoteRune(elem))
	}

	return elems
}

// RunesToStrings returns a slice of strings from a slice of runes.
//
// Parameters:
//   - slice: The slice of runes.
//
// Returns:
//   - []string: The slice of strings. Nil if the slice is empty.
//
// Deprecated: Function moved over the strings package of mygo-lib.
func RunesToStrings(slice []rune) []string {
	if len(slice) == 0 {
		return nil
	}

	elems := make([]string, 0, len(slice))

	for _, elem := range slice {
		elems = append(elems, string(elem))
	}

	return elems
}

// EitherOrString is a function that returns a string representation of a slice
// of strings. Empty strings are ignored.
//
// Parameters:
//   - values: The values to convert to a string.
//
// Returns:
//   - string: The string representation.
//
// Example:
//
//	EitherOrString([]string{"a", "b", "c"}) // "either a, b, or c"
func EitherOrString(elems []string) string {
	if len(elems) == 0 {
		return ""
	}

	if len(elems) == 1 {
		return elems[0]
	}

	var builder strings.Builder

	builder.WriteString("either ")

	if len(elems) > 2 {
		builder.WriteString(strings.Join(elems[:len(elems)-1], ", "))
		builder.WriteRune(',')
	} else {
		builder.WriteString(elems[0])
	}

	builder.WriteString(" or ")
	builder.WriteString(elems[len(elems)-1])

	return builder.String()
}
