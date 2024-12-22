package common

// Fixer is an interface for objects that can fix their internal state.
type Fixer interface {
	// Fix fixes the internal state of the object.
	//
	// Returns:
	//   - error: If the object could not be fixed.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error: Implementation-specific.
	Fix() error
}

// Fix attempts to fix the internal state of the given Fixer object.
//
// Parameters:
//   - fixer: The Fixer object whose internal state is to be fixed.
//   - kind: A string representing the kind of object being fixed. If empty, defaults to "object".
//
// Returns:
//   - error: An error if the object could not be fixed, or if the provided object is nil.
//
// Errors:
//   - *ErrBadParam: If obj is nil.
//   - *ErrWhile: If the object could not be fixed.
func Fix(fixer Fixer, kind string) error {
	if kind == "" {
		kind = "object"
	}

	var err error

	if fixer == nil {
		err = NewErrNilParam(kind)
	} else {
		err = fixer.Fix()
		if err == nil {
			return nil
		}
	}

	err = NewErrWhile("fixing "+kind, err)
	return err
}
