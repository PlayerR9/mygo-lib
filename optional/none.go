package optional

// noneOptional is an Optional that has no value.
type noneOptional struct{}

// IsPresent implements Optional.
//
// Always returns false.
func (no noneOptional) IsPresent() bool {
	return false
}

// Get implements Optional.
//
// Always returns an error.
func (no noneOptional) Get() (any, error) {
	return nil, ErrMissingValue
}

// None creates an Optional with no value.
//
// Returns:
//   - Optional: An Optional with no value. Never returns nil.
func None() Optional {
	return noneOptional{}
}
