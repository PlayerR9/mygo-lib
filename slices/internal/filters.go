package internal

// Filter filters a slice of E's according to a predicate.
//
// Parameters:
//   - s: The slice of E's to filter.
//   - predicate: The predicate to use for filtering.
//
// Returns:
//   - uint: The number of elements rejected.
func Filter[S ~[]E, E any](s *S, predicate func(e E) bool) uint {
	var end uint

	for _, v := range *s {
		ok := predicate(v)
		if !ok {
			continue
		}

		(*s)[end] = v
		end++
	}

	removed := uint(len(*s)) - end

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return removed
}

// Reject rejects a slice of E's according to a predicate.
//
// Parameters:
//   - s: The slice of E's to reject.
//   - predicate: The predicate to use for rejecting.
//
// Returns:
//   - uint: The number of elements rejected.
func Reject[S ~[]E, E any](s *S, predicate func(e E) bool) uint {
	var end uint

	for _, v := range *s {
		ok := predicate(v)
		if ok {
			continue
		}

		(*s)[end] = v
		end++
	}

	rejected := uint(len(*s)) - end

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return rejected
}

// RejectNils rejects in-place all nil elements from a slice of E's and returns the number
// of elements rejected.
//
// Parameters:
//   - s: The slice of E's to filter.
//
// Returns:
//   - uint: The number of elements rejected.
func RejectNils[S ~[]*E, E any](s *S) uint {
	lenS := uint(len(*s))

	var end uint
	for end < lenS && (*s)[end] != nil {
		end++
	}

	for i := uint(0); i < end; i++ {
		if (*s)[i] != nil {
			continue
		}

		(*s)[i] = (*s)[end-1]
		(*s)[end-1] = nil
		end--

		i-- // retry since we have overwritten the element
	}

	rejected := lenS - end

	clear((*s)[end:])
	*s = (*s)[:end]

	return rejected
}

// RejectZero rejects in-place all zero-valued elements from a slice of E's and returns the number
// of elements rejected.
//
// Parameters:
//   - s: The slice of E's to filter.
//
// Returns:
//   - uint: The number of elements rejected.
func RejectZero[S ~[]E, E comparable](s *S) uint {
	zero := *new(E)

	var end uint

	for _, v := range *s {
		if v == zero {
			continue
		}

		(*s)[end] = v
		end++
	}

	rejected := uint(len(*s)) - end

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return rejected
}
