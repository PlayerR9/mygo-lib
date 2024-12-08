package common

// ClearFrom clears all elements in the slice starting from the given index. If
// the given index is equal to 0, the entire slice is cleared. If the
// given index is greater than or equal to the length of the slice, the function
// does nothing.
//
// Parameters:
//   - slice: The slice to clear elements from.
//   - from_idx: The index to start clearing elements from.
//
// Panics:
//   - ErrBadParam: If the slice is nil.
func ClearFrom[T any](slice *[]T, from_idx uint) {
	if slice == nil {
		panic(NewErrNilParam("slice"))
	}

	if from_idx >= uint(len(*slice)) {
		return
	}

	if from_idx == 0 {
		clear(*slice)
		*slice = nil

		return
	}

	clear((*slice)[from_idx:])
	*slice = (*slice)[:from_idx]
}

// ClearTo clears all elements in the slice up to (but not including) the given
// index. If the given index is  0, the function does nothing. If the
// given index is greater than or equal to the length of the slice, the entire
// slice is cleared.
//
// Parameters:
//   - slice: The slice to clear elements from.
//   - to_idx: The index up to which to clear elements.
func ClearTo[T any](slice *[]T, to_idx uint) {
	if slice == nil {
		panic(NewErrNilParam("slice"))
	}

	if to_idx == 0 {
		return
	}

	if to_idx >= uint(len(*slice)) {
		clear(*slice)
		*slice = nil

		return
	}

	clear((*slice)[:to_idx])
	*slice = (*slice)[to_idx:]
}
