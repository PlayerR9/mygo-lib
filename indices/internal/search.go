package internal

// FirstIndexOf returns the first index of the slice for which the predicate returns true, or an empty index
// if no element satisfies the predicate.
//
// Parameters:
//   - s: The slice to search.
//   - predicate: The predicate to use when searching.
//
// Returns:
//   - uint: The first index of the slice that satsifies the predicate.
//   - bool: True if an element satisfies the predicate, false otherwise.
func FirstIndexOf[S ~[]E, E any](s S, predicate func(e E) bool) (uint, bool) {
	for i, elem := range s {
		ok := predicate(elem)
		if ok {
			return uint(i), true
		}
	}

	return 0, false
}
