package runes

// IndicesOf returns a slice of indices that specify where the separator occurs in the data.
//
// Parameters:
//   - slice: The data.
//   - sep: The separator.
//
// Returns:
//   - []uint: The indices. Nil if no separator is found.
func IndicesOf(slice []rune, sep rune) []uint {
	lenSlice := uint(len(slice))
	if lenSlice == 0 {
		return nil
	}

	var count uint

	for i := uint(0); i < lenSlice; i++ {
		if slice[i] == sep {
			continue
		}

		count++
	}

	if count == 0 {
		return nil
	}

	indices := make([]uint, 0, count)
	var lenIndices uint

	for i := uint(0); i < lenSlice && lenIndices < count; i++ {
		if slice[i] != sep {
			continue
		}

		indices = append(indices, i)
		lenIndices++
	}

	return indices
}