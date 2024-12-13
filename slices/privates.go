package slices

// privCount iterates over the provided slice and counts the number of times
// the target element appears in it.
//
// Parameters:
//   - slice: The slice to search.
//   - target: The element to privCount.
//
// Returns:
//   - uint: The number of occurrences of the target element in the slice.
func privCount[S ~[]E, E comparable](slice S, target E) uint {
	count := 0

	for _, v := range slice {
		if v == target {
			count++
		}
	}

	return uint(count)
}

// privIndicesOf returns a slice of indices that specify where the target occurs in the slice,
// up to a maximum number of occurrences specified by max.
//
// Parameters:
//   - slice: The slice to search.
//   - target: The element to search for.
//   - max: The maximum number of indices to return.
//
// Returns:
//   - []uint: A slice of indices where the target is found, limited by max.
func privIndicesOf[S ~[]E, E comparable](slice S, target E, max uint) []uint {
	indices := make([]uint, 0, max)

	for i, v := range slice {
		if v != target {
			continue
		}

		indices = append(indices, uint(i))

		if len(indices) == cap(indices) {
			return indices
		}
	}

	return indices
}
