package slices

import "github.com/PlayerR9/mygo-lib/optional"

// Count returns the number of times the target appears in the slice.
//
// Parameters:
//   - slice: The slice to search.
//   - target: The target to count.
//
// Returns:
//   - uint: The number of occurrences of the target in the slice.
func Count[S ~[]E, E comparable](slice S, target E) uint {
	if len(slice) == 0 {
		return 0
	}

	n := privCount(slice, target)
	return n
}

// IndicesOf returns a slice of indices where the target occurs in the slice
// up to a maximum number of occurrences specified by max.
//
// Parameters:
//   - slice: The slice to search.
//   - target: The target to search for.
//   - max: The maximum number of indices to return.
//
// Returns:
//   - []uint: A slice of indices where the target is found, limited by max.
func IndicesOf[S ~[]E, E comparable](slice S, target E, max uint) []uint {
	if max == 0 || len(slice) == 0 {
		return nil
	}

	indices := privIndicesOf(slice, target, max)
	return indices
}

// AllIndicesOf is a convenience function that returns a slice of indices where the target
// occurs in the slice.
//
// Parameters:
//   - slice: The slice to search.
//   - target: The target to search for.
//
// Returns:
//   - []uint: A slice of indices where the target is found.
func AllIndicesOf[S ~[]E, E comparable](slice S, target E) []uint {
	max := privCount(slice, target)
	if max == 0 {
		return nil
	}

	indices := privIndicesOf(slice, target, max)
	return indices
}

// FirstIndexOf returns the first index of the target in the slice and a boolean indicating
// whether the target was found.
//
// Parameters:
//   - slice: The slice to search.
//   - target: The target to search for.
//
// Returns:
//   - optional.Optional: The index of the target if found, or None if not found.
//
// The optional is of type uint when it exists.
func FirstIndexOf[S ~[]E, E comparable](slice S, target E) optional.Optional {
	if len(slice) == 0 {
		return optional.None()
	}

	for i, v := range slice {
		if v == target {
			return optional.Some(uint(i))
		}
	}

	return optional.None()
}

// LastIndexOf returns the last index of the target in the slice and a boolean indicating
// whether the target was found.
//
// Parameters:
//   - slice: The slice to search.
//   - target: The target to search for.
//
// Returns:
//   - optional.Optional: The index of the target if found, or None if not found.
//
// The optional is of type uint when it exists.
func LastIndexOf[S ~[]E, E comparable](slice S, target E) optional.Optional {
	if len(slice) == 0 {
		return optional.None()
	}

	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return optional.Some(uint(i))
		}
	}

	return optional.None()
}
