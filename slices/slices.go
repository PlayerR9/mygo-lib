package slices

import (
	"cmp"
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
)

// MayInsert attempts to insert an element into a sorted slice if it is not already present.
//
// Parameters:
//   - slice: A pointer to a slice of ordered elements.
//   - elem: The element to insert.
//
// Returns:
//   - bool: Returns true if the element was inserted into the slice, false otherwise.
//
// If the element is not found in the slice, it is inserted in the correct position to maintain order.
//
// Panics:
//   - common.ErrBadParam: If slice is nil.
func MayInsert[T cmp.Ordered](slice *[]T, elem T) bool {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	pos, ok := slices.BinarySearch(*slice, elem)
	if ok {
		return false
	}

	*slice = slices.Insert(*slice, pos, elem)

	return true
}

// Uniquefy removes duplicate elements from a slice, in-place, while also sorting the slice in
// ascending order.
//
// Parameters:
//   - slice: A pointer to the slice where duplicate elements will be removed from.
//
// Returns:
//   - uint: The number of elements removed from the slice.
//
// The slice is also resized while clearing the removed elements.
//
// Panics:
//   - common.ErrBadParam: If slice is nil.
func Uniquefy[T cmp.Ordered](slice *[]T) uint {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice < 2 {
		return 0
	}

	limit := uint(1)

	for i := uint(1); i < lenSlice; i++ {
		elem := (*slice)[i]

		pos, ok := slices.BinarySearch((*slice)[:limit], elem)
		if ok {
			continue
		}

		for j := limit; j >= uint(pos); j-- {
			(*slice)[j] = (*slice)[j-1]
		}

		(*slice)[pos] = elem
		limit++
	}

	clear((*slice)[limit:])

	n := lenSlice - limit

	*slice = (*slice)[:limit]

	return n
}

// Merge inserts elements from the 'from' slice into the 'dest' slice, maintaining order and ensuring no duplicates.
//
// Parameters:
//   - dest: A pointer to the destination slice where elements will be inserted. This slice must be sorted
//     and free of duplicates. If not, `Uniquefy()` must be called on it first.
//   - from: The slice of elements to merge into the destination.
//
// Returns:
//   - uint: The number of elements that were not inserted or could not be merged.
//
// If 'from' is empty, the function does nothing. Each element from 'from' is inserted into 'dest' in the correct position,
// ensuring that 'dest' remains sorted and free of duplicates.
//
// Panics:
//   - common.ErrBadParam: If dest is nil.
func Merge[T cmp.Ordered](dest *[]T, from []T) uint {
	if len(from) == 0 {
		return 0
	} else if dest == nil {
		panic(common.NewErrNilParam("dest"))
	}

	var n uint

	for _, elem := range from {
		pos, ok := slices.BinarySearch(*dest, elem)
		if ok {
			n++
		} else {
			*dest = slices.Insert(*dest, pos, elem)
		}
	}

	return n
}
