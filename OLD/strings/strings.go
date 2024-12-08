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

// OrQuoteElse returns the quoted string if the string is not empty. Otherwise, it returns the default string.
//
// Parameters:
//   - str: The string to quote if it is not empty.
//   - def: The default string to return if str is empty.
//
// Returns:
//   - string: The quoted string or the default string.
func OrQuoteElse(str string, def string) string {
	if str == "" {
		return def
	} else {
		return strconv.Quote(str)
	}
}
