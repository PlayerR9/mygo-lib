package strings

import (
	"strconv"
)

// Quote quotes each string in the given slice of strings in-place.
//
// Parameters:
//   - elems: The slice of strings to quote.
//
// Example:
//
//	Quote([]string{"a", "b", "c"}) // => []string{"\"a\"", "\"b\"", "\"c\""}
func Quote(elems []string) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return
	}

	for i := uint(0); i < lenElems; i++ {
		elems[i] = strconv.Quote(elems[i])
	}
}
