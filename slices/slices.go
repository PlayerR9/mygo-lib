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
//   - error: Returns ErrBadParam if slice is nil.
//
// If the element is not found in the slice, it is inserted in the correct position to maintain order.
func MayInsert[T cmp.Ordered](slice *[]T, elem T) error {
	if slice == nil {
		return common.NewErrNilParam("slice")
	}

	pos, ok := slices.BinarySearch(*slice, elem)
	if ok {
		return nil
	}

	*slice = slices.Insert(*slice, pos, elem)

	return nil
}

// Merge inserts elements from the 'from' slice into the 'dest' slice, maintaining order and ensuring no duplicates.
//
// Parameters:
//   - dest: A pointer to the destination slice where elements will be inserted.
//   - from: The slice of elements to merge into the destination.
//
// Returns:
//   - error: Returns ErrBadParam if dest is nil.
//
// If 'from' is empty, the function does nothing. Each element from 'from' is inserted into 'dest' in the correct position,
// ensuring that 'dest' remains sorted and free of duplicates.
func Merge[T cmp.Ordered](dest *[]T, from []T) error {
	if len(from) == 0 {
		return nil
	} else if dest == nil {
		return common.NewErrNilParam("dest")
	}

	for _, elem := range from {
		pos, ok := slices.BinarySearch(*dest, elem)
		if ok {
			continue
		}

		*dest = slices.Insert(*dest, pos, elem)
	}

	return nil
}
