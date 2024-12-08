package mem

// Releaseable is an interface defining the behavior of types that can
// release their memory.
type Releaseable interface {
	// Release releases ONLY the memory used by the type; thus, later use
	// of the type are not guaranteed to succeed (i.e., fail or undefined
	// behavior). This also implies that, despite not hindering thread-safety,
	// goroutines might fail once `Release()` is called anywhere.
	//
	// Therefore, in order to guarantee thread-safety, programmers should
	// handle the case where the type is released outside of their control.
	//
	// Finally, most of the time, this is used in conjunction with defer
	// statements to ensure that the memory is released as soon as possible.
	//
	// WARNING: This must NOT be called directly. Instead, use the provided
	// `Release()` function.
	//
	// Returns:
	//   - error: An error that may occur when releasing the memory.
	//
	// Errors:
	//   - ErrNilReceiver: If the receiver is nil.
	//   - ErrInvalidObject: If the method `Release()` is called already.
	//   - any other error returned by the `Release()` method.
	Release() error
}

// Release calls the `Release()` method on an interface that implements Releaseable and releases the memory used by it.
//
// Parameters:
//   - target: The name of the target whose memory is being released.
//   - arg: A pointer to a Releaseable interface implementation.
//
// Returns:
//   - error: An error that may occur when releasing the memory.
//
// Errors:
//   - NewErrNilParam: If the argument or its dereferenced value is nil.
//   - *ErrRelease: If the `Release()` method returns an error.
func Release(target string, arg Releaseable) error {
	if arg == nil {
		return NewErrNilParam("arg")
	}

	err := arg.Release()
	if err == nil {
		return nil
	}

	e, ok := err.(*ErrRelease)
	if !ok {
		return NewErrRelease(target, err)
	}

	_ = e.AppendTarget(target)

	return e
}
