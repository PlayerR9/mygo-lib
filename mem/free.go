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
	//
	// Returns:
	//   - error: An error that may occur when releasing the memory.
	//
	// Errors:
	//   - ErrNilReceiver: If the receiver is nil.
	//   - ErrInvalidObject: If the method `Free()` is called already.
	//   - any other error returned by the `Free()` method.
	Free() error
}

// Free calls the `Free()` method on a Freeable interface and releases the memory used by it.
//
// Parameters:
//   - target: The name of the target whose memory is being released.
//   - arg: A pointer to a Freeable interface implementation.
//
// Returns:
//   - error: An error that may occur when releasing the memory.
//
// Errors:
//   - NewErrNilParam: If the argument or its dereferenced value is nil.
//   - *ErrFree: If the `Free()` method returns an error.
func Free(target string, arg Freeable) error {
	if arg == nil {
		return NewErrNilParam("arg")
	}

	err := arg.Free()
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

// FreeUnless calls the `Free()` method on a Freeable interface and releases the memory used by it,
// unless the `Free()` method returns ErrNilReceiver or ErrInvalidObject.
//
// Parameters:
//   - target: The name of the target whose memory is being released.
//   - arg: A pointer to a Freeable interface implementation.
//
// Returns:
//   - first bool: true if the memory was released, false otherwise.
//   - second error: An error that may occur when releasing the memory, if any.
//
// Errors:
//   - NewErrNilParam: If the argument or its dereferenced value is nil.
//   - *ErrFree: If the `Free()` method returns an error and it is not ErrNilReceiver or ErrInvalidObject.
//
// In essence, this is used for optionally releasing memory if not done already. Therefore, this is a
// more comprehensive shorthand for the following:
//
//	if arg != nil {
//		if err := Free("arg", arg); err != nil {
//			return err
//		}
//	}
func FreeUnless(target string, arg Freeable) (bool, error) {
	if arg == nil {
		return false, nil
	}

	err := arg.Free()
	if err == nil || err == ErrNilReceiver {
		return true, nil
	}

	switch err := err.(type) {
	case *ErrInvalidObject:
		return false, nil
	case *ErrRelease:
		_ = err.AppendTarget(target)
		return false, err
	default:
		return false, NewErrRelease(target, err)
	}
}
