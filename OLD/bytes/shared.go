package bytes

// IndicesOf returns a slice of indices that specify where the separator occurs in the data.
//
// Parameters:
//   - slice: The data.
//   - sep: The separator.
//
// Returns:
//   - []int: The indices. Nil if no separator is found.
func IndicesOf(slice []byte, sep []byte) []uint {
	lenSep := uint(len(sep))
	if lenSep == 0 {
		return nil
	}

	lenSlice := uint(len(slice))
	if lenSlice == 0 {
		return nil
	}

	var count uint

	for i := uint(0); i < lenSlice; i++ {
		if slice[i] != sep[0] {
			continue
		}

		count++
	}

	if count == 0 {
		return nil
	}

	indices := make([]uint, 0, count)

	for i := uint(0); i < lenSlice-lenSep+1; i++ {
		if slice[i] != sep[0] {
			continue
		}

		indices = append(indices, i)
	}

	var top uint

	for i := uint(1); i < lenSep; i++ {
		top = 0

		for _, idx := range indices {
			if slice[idx+1] != sep[i] {
				continue
			}

			indices[top] = idx
			top++
		}

		if top == 0 {
			return nil
		}

		indices = indices[:top]
	}

	indices = indices[:len(indices):len(indices)]

	return indices
}
