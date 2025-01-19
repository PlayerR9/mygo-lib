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
func Quote(s []string) {
	if len(s) == 0 {
		return
	}

	for i, e := range s {
		str := strconv.Quote(e)
		s[i] = str
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

	switch len(s) {
	case 0:
		return ""
	case 1:
		return s[0]
	case 2:
		return "either " + s[0] + " or " + s[1]
	default:
		var builder strings.Builder

		_, _ = builder.WriteString("either ")
		_, _ = builder.WriteString(s[0])

		for _, elem := range s[1 : len(s)-1] {
			_, _ = builder.WriteString(", ")
			_, _ = builder.WriteString(elem)
		}

		_, _ = builder.WriteRune(',')
		_, _ = builder.WriteString(s[len(s)-1])

		str := builder.String()
		return str
	}

}
