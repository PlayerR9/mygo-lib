package pointers

import "github.com/PlayerR9/mygo-lib/common"

// Pointer is an interface for pointer-like types.
type Pointer interface {
	// IsNil checks whether the pointer is nil.
	//
	// Returns:
	//   - bool: True if the pointer is nil, false otherwise.
	IsNil() bool
}

// RejectNils removes in-place all nil elements from the given slice of pointer-like
// types that implement the Pointer interface.
//
// Parameters:
//   - slice: The slice of pointer-like types to remove nils from.
//
// Returns:
//   - uint: The number of elements removed from the slice.
//
// Panics:
//   - common.ErrBadParam: If the slice is nil.
func RejectNils[T Pointer](slice *[]T) uint {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0
	}

	var top uint

	for _, elem := range *slice {
		ok := elem.IsNil()
		if ok {
			continue
		}

		(*slice)[top] = elem
		top++
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]

	return lenSlice - top
}
