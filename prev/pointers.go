package common

// Niler is a type that can be used to check if a pointer is nil.
type Niler interface {
	// IsNil checks whether the pointer is nil.
	//
	// Returns:
	//   - bool: True if the pointer is nil, false otherwise.
	IsNil() bool
}

// Set sets the value of the pointer if it is not nil.
//
// Parameters:
//   - p: The pointer to set.
//   - v: The value to set the pointer to.
func Set[T any](p *T, v T) {
	if p == nil {
		return
	}

	*p = v
}

// Get returns the value of the pointer if it is not nil.
//
// Parameters:
//   - p: The pointer to get.
//
// Returns:
//   - T: The value of the pointer.
//   - bool: True if the pointer is not nil, false otherwise.
func Get[T any](p *T) T {
	if p == nil {
		return *new(T)
	}

	return *p
}
