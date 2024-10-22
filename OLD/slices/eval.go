package slices

// ComplexFilter applies a filter function on a slice of elements based on the provided filter function.
//
// This function uses indices for optimization reasons.
//
// Parameters:
//   - slice: The slice of elements to filter.
//   - fn: The filter function that takes a list of indices and returns a list of selected indices and a
//     flag indicating early exit.
//
// Returns:
//   - []T: The filtered slice of elements.
//
// Behavior:
//   - If the provided slice is empty or the filter function is nil, returns nil.
//   - The filter function should return the selected indices and a flag to indicate early exit.
//   - The filtered slice contains only the elements corresponding to the selected indices.
func ComplexFilter[T any](slice []T, fn func(indices []int) ([]int, bool)) []T {
	if len(slice) == 0 || fn == nil {
		return nil
	}

	indices := make([]int, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		indices = append(indices, i)
	}

	var early_exit bool

	for len(indices) > 0 && !early_exit {
		indices, early_exit = fn(indices)
	}

	if !early_exit {
		return nil
	}

	remaining := make([]T, 0, len(indices))

	for _, idx := range indices {
		remaining = append(remaining, slice[idx])
	}

	return remaining
}
