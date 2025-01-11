package internal

import (
	"strconv"
	"strings"
)

// RejectEmpty rejects all empty strings from a slice of strings and returns a
// new slice containing the filtered strings.
//
// Parameters:
//   - s: The slice of strings to filter.
//
// Returns:
//   - []string: A new slice of strings with all empty strings removed.
func RejectEmpty(s []string) []string {
	var count uint

	for _, str := range s {
		if str != "" {
			count++
		}
	}

	if count == 0 {
		return nil
	}

	result := make([]string, 0, count)

	for _, str := range s {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}

// Quote quotes all the strings in the given slice with strconv.Quote.
//
// Parameters:
//   - s: The slice of strings to quote.
//
// Example:
//
//	strs := []string{"hello", " world"}
//	Quote(strs)
//	fmt.Println(strs) // prints ["hello", " world"]
func Quote(s []string) {
	for i := range s {
		s[i] = strconv.Quote(s[i])
	}
}

// EitherOr returns a string that says either this or that, or a list of
// options.
//
// Parameters:
//   - s: The slice of strings to choose from.
//
// Returns:
//   - string: The string that says either this or that.
//
// Example:
//
//	EitherOr([]string{"a", "b"}) // returns "either a or b"
//	EitherOr([]string{"a", "b", "c"}) // returns "either a, b, or c"
func EitherOr(s []string) string {
	var builder strings.Builder

	if len(s) > 2 {
		_, _ = builder.WriteString("either ")
	}

	_, _ = builder.WriteString(s[0])

	if len(s) > 2 {
		for _, str := range s[1 : len(s)-1] {
			_, _ = builder.WriteString(", ")
			_, _ = builder.WriteString(str)
		}

		_, _ = builder.WriteRune(',')
	}

	if len(s) > 1 {
		_, _ = builder.WriteString(" or ")
		_, _ = builder.WriteString(s[len(s)-1])
	}

	str := builder.String()
	return str
}
