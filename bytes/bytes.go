package bytes

// indicesOfByte returns a slice of indices that specify where the separator occurs in the data.
//
// Parameters:
//   - data: The data.
//   - sep: The separator.
//   - size: The size of the separator.
//
// Returns:
//   - []int: The indices.
func indicesOfByte(data []byte, sep byte, size int) []int {
	size = len(data) - size

	var count int

	for i := 0; i < size; i++ {
		if data[i] == sep {
			count++
		}
	}

	if count == 0 {
		return nil
	}

	indices := make([]int, 0, count)

	for i := 0; i < size; i++ {
		if data[i] == sep {
			indices = append(indices, i)
		}
	}

	return indices
}

// IndicesOf returns a slice of indices that specify where the separator occurs in the data.
//
// Parameters:
//   - data: The data.
//   - sep: The separator.
//
// Returns:
//   - []int: The indices.
//
// WARNING: This function is yet to be tested. However, it should work as expected.
func IndicesOf(data []byte, sep []byte) []int {
	sep_len := len(sep)

	if len(data) == 0 || sep_len == 0 {
		return nil
	}

	b := sep[0]

	indices := indicesOfByte(data, b, sep_len)
	if len(indices) == 0 {
		return nil
	}

	limit := len(indices)

	for i := 1; i < sep_len; i++ {
		b := sep[i]

		var top int

		for j := 0; j < limit; j++ {
			idx := indices[j] + i

			if data[idx] == b {
				indices[top] = idx
				top++
			}
		}

		if top == 0 {
			return nil
		}

		limit = top
	}

	return indices[:limit:limit]
}
