package slices

import "github.com/PlayerR9/mygo-lib/optional"

// IndicesOf returns a slice of indices that specify where the target occurs in the slice.
//
// Parameters:
//   - slice: The slice to search.
//   - target: The target to search for.
//
// Returns:
//   - []uint: The indices.
func IndicesOf[S ~[]E, E comparable](slice S, target E) []uint {
	if len(slice) == 0 {
		return nil
	}

	var indices []uint

	for i, v := range slice {
		if v == target {
			indices = append(indices, uint(i))
		}
	}

	return indices
}

func IndicesOfMax[]

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
