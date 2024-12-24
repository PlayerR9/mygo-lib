package indices

// MustGet returns the value of the Index if it is Some, or panics if it is None.
//
// This function is useful when you are sure that the Index is Some, and you want
// to get the value without having to use an "if" statement.
//
// Parameters:
//   - idx_opt: The Index to get the value from.
//
// Returns:
//   - uint: The value of the Index.
func MustGet(idx_opt Index) uint {
	idx, err := idx_opt.Get()
	if err != nil {
		panic(err)
	}

	return idx
}
