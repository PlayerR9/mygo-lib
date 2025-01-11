package runes

// Of returns a predicate function that checks if a given rune matches
// the specified rune.
//
// Parameters:
//   - r: The rune to match against.
//
// Returns:
//   - func(rune) bool: A predicate function that returns true if the
//     input rune matches the specified rune, false otherwise. (never nil)
func Of(r rune) func(rune) bool {
	predicate := func(char rune) bool {
		return char == r
	}

	return predicate
}
