package common

// Resetter is an interface defining the behavior of types that can
// reset their internal state.
type Resetter interface {
	// Reset clears the type's internal state in order to make it ready
	// for reuse.
	//
	// Unlike `Free()`, this method guarantees that the type is in a
	// valid state.
	Reset()
}

// Reset calls the `Reset()` method on the given argument if it implements the
// `Resetter` interface. If the argument is `nil` or does not implement
// `Resetter`, this function does nothing.
//
// This function is intended to be used to reset the internal state of a given
// type so that it can be reused.
//
// Returns:
//   - bool: True if the type was resetted, false otherwise.
func Reset(arg any) bool {
	if arg == nil {
		return true
	}

	a, ok := arg.(interface{ Reset() })
	if !ok {
		return false
	}

	a.Reset()

	return true
}
