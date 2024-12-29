package indices

import "github.com/PlayerR9/mygo-lib/indices/internal"

// Predicate is a function that returns true if the element is valid.
//
// Parameters:
//   - e: The element to check.
//
// Returns:
//   - bool: True if the element is valid, false otherwise.
type Predicate[E any] func(e E) bool

// FirstIndexOf returns the first index of the slice for which the predicate returns true, or an empty index
// if no element satisfies the predicate.
//
// Parameters:
//   - s: The slice to search.
//   - predicate: The predicate to use when searching.
//
// Returns:
//   - Index: The first index of the slice that satsifies the predicate, or an empty index if none do.
func FirstIndexOf[S ~[]E, E any](s S, predicate Predicate[E]) Index {
	if len(s) == 0 || predicate == nil {
		return None()
	}

	idx, ok := internal.FirstIndexOf(s, predicate)
	if ok {
		return Some(idx)
	} else {
		return None()
	}
}
