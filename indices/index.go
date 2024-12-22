package indices

// Index is a struct that holds an optional index.
type Index struct {
	// idx is the index.
	idx *uint
}

// Some creates an Index with the given value.
//
// Parameters:
//   - idx: The index to wrap in the Index.
//
// Returns:
//   - Index: The newly created Index.
func Some(idx uint) Index {
	opt := Index{
		idx: &idx,
	}

	return opt
}

// None creates an Index with no value.
//
// Returns:
//   - Index: The newly created empty Index.
func None() Index {
	opt := Index{}
	return opt
}

// OrElse returns the value of the index if it is not empty, otherwise it returns the fallback value.
//
// Parameters:
//   - fallback: The value to return if the index is empty.
//
// Returns:
//   - uint: The value of the index if not empty, otherwise fallback.
func (idx Index) OrElse(fallback uint) uint {
	if idx.idx == nil {
		return fallback
	} else {
		return *idx.idx
	}
}
