package common

// Make creates a slice of type T with the specified capacity.
//
// Parameters:
//   - n: The capacity of the slice to create. If zero, returns nil.
//
// Returns:
//   - []T: A newly created slice with the specified capacity.
func Make[T any](n uint32) []T {
	if n == 0 {
		return nil
	}

	slice := make([]T, 0, n)
	return slice
}
