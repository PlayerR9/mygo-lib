package strings

import (
	"strconv"
)

// Quote quotes each string in the given slice of strings.
//
// Parameters:
//   - elems: The slice of strings to quote.
//
// Example:
//
//	Quote([]string{"a", "b", "c"}) // => []string{"\"a\"", "\"b\"", "\"c\""}
func Quote(elems []string) {
	for i := range elems {
		elems[i] = strconv.Quote(elems[i])
	}
}
