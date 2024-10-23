package runes

/*
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
*/
