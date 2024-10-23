package pointers

// Pointer is an interface for pointer-like types.
type Pointer interface {
	// IsNil checks whether the pointer is nil.
	//
	// Returns:
	//   - bool: True if the pointer is nil, false otherwise.
	IsNil() bool
}

// RejectNils removes all nil elements from the given slice of pointer-like
// types that implement the Pointer interface.
//
// Parameters:
//   - slice: The slice of pointer-like types to remove nils from.
//
// Returns:
//   - []T: The slice of pointer-like types without nils. Nil if all the
//     elements are nil or no elements were specified.
func RejectNils[T Pointer](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	var count int

	for i := range slice {
		ok := slice[i].IsNil()
		if !ok {
			count++
		}
	}

	if count == 0 {
		return nil
	}

	new_slice := make([]T, 0, count)

	for i := range slice {
		ok := slice[i].IsNil()
		if !ok {
			new_slice = append(new_slice, slice[i])
		}
	}

	return new_slice
}
