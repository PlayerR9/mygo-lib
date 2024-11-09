package common

// Freeable is an interface defining the behavior of types that can
// release their memory.
type Freeable interface {
	// Free releases ONLY the memory used by the type; thus, later use
	// of the type are not guaranteed to succeed (i.e., fail or undefined
	// behavior). This also implies that, despite not hindering thread-safety,
	// goroutines might fail once `Free()` is called anywhere.
	//
	// Therefore, in order to guarantee thread-safety, programmers should
	// handle the case where the type is freed outside of their control.
	//
	// Finally, most of the time, this is used in conjunction with defer
	// statements to ensure that the memory is released as soon as possible.
	Free()
}

// Free calls the `Free()` method on the given argument if it implements the
// `Type` interface. If the argument is `nil` or does not implement `Type`,
// this function does nothing.
//
// Returns:
//   - bool: True if the type was freed, false otherwise.
func Free(arg any) bool {
	if arg == nil {
		return true
	}

	a, ok := arg.(interface{ Free() })
	if !ok {
		return false
	}

	a.Free()

	return true
}

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
