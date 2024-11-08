package common

// Type is an interface defining the behavior of most types.
type Type interface {
	// Free releases the memory used by the type and, as such, successive
	// use of the type after calling `Free()` may fail or cause undefined
	// behavior. As such, despite being a thread-safe method, goroutines
	// that are using the type may fail once `Free()` is called.
	//
	// In such cases, it's up to the programmer to keep track of when a
	// given type is no longer in use.
	//
	// Most of the time, this is used in conjunction with `defer` to ensure
	// that the memory is released as soon as possible.
	Free()
}

// Free calls the `Free()` method on the given argument if it implements the
// `Type` interface. If the argument is `nil` or does not implement `Type`,
// this function does nothing.
func Free(arg any) {
	if arg == nil {
		return
	}

	a, ok := arg.(interface{ Free() })
	if !ok {
		return
	}

	a.Free()
}
