package mem

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
	//
	// WARNING: This must NOT be called directly. Instead, use the provided
	// `Free()` function.
	Free()
}

// Free calls the `Free()` method on the Freeable stored in the pointer and sets it to nil.
// This is a wrapper function that is thread-safe.
//
// Parameters:
//   - arg: A double pointer to the Ref instance to be released.
func Free(arg **Ref) {
	if arg == nil || *arg == nil {
		return
	}

	(*arg).free()
	(*arg).free = nil
	(*arg).ptr = nil

	*arg = nil
}
