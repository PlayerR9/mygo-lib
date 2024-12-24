package indices

const (
	// MaxUint is the maximum value of a uint.
	MaxUint uint = ^uint(0)
)

// Index is a type alias for uint. MaxUint is used to represent an empty index.
type Index uint

// Some creates an Index with the given value.
//
// Parameters:
//   - idx: The index to wrap in the Index.
//
// Returns:
//   - Index: The newly created Index.
func Some(idx uint) Index {
	opt := Index(idx)

	return opt
}

// None creates an Index with no value.
//
// Returns:
//   - Index: The newly created empty Index.
func None() Index {
	opt := Index(MaxUint)
	return opt
}

// IsPresent checks if the index is not empty.
//
// Returns:
//   - bool: True if the index is not empty, false otherwise.
func (idx Index) IsPresent() bool {
	return uint(idx) != MaxUint
}

// Get retrieves the underlying uint value of the Index if it is present.
//
// Returns:
//   - uint: The underlying value of the Index.
//   - error: An error if the Index is empty.
//
// Errors:
//   - ErrMissingValue: If the Index is empty.
func (idx Index) Get() (uint, error) {
	underlying := uint(idx)
	if underlying == MaxUint {
		return 0, ErrMissingValue
	}

	return underlying, nil
}

// OrElse returns the value of the index if it is not empty, otherwise it returns the fallback value.
//
// Parameters:
//   - fallback: The value to return if the index is empty.
//
// Returns:
//   - uint: The value of the index if not empty, otherwise fallback.
func (idx Index) OrElse(fallback uint) uint {
	underlying := uint(idx)

	if underlying != MaxUint {
		fallback = underlying
	}

	return fallback
}
