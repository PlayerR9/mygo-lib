package strings

import (
	"strconv"
	"strings"
)

// Quote quotes all the strings in the given slice with strconv.Quote.
//
// Parameters:
//   - str: The slice of strings to quote.
//
// Example:
//
//	strs := []string{"hello", " world"}
//	strings.Quote(strs)
//	fmt.Println(strs) // prints ["hello", " world"]
func Quote(str []string) {
	if len(str) == 0 {
		return
	}

	for i := range str {
		str[i] = strconv.Quote(str[i])
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
