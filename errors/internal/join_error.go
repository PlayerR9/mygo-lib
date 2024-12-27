package internal

import "bytes"

// joinError is an error that joins multiple errors into a single string.
type joinError struct {
	// errs is the slice of errors to join.
	errs []error
}

// Error implements error.
func (je joinError) Error() string {
	elems := make([][]byte, 0, len(je.errs))

	for _, err := range je.errs {
		msg := err.Error()
		if msg == "" {
			continue
		}

		elems = append(elems, []byte(msg))
	}

	if len(elems) == 0 {
		return ""
	}

	data := bytes.Join(elems, []byte{'\n'})

	return string(data)
}

// NewJoinError creates a new joinError instance that combines multiple errors.
//
// Parameters:
//   - errs: A slice of errors to be joined.
//
// Returns:
//   - error: An instance of joinError. Never returns nil.
func NewJoinError(errs []error) error {
	je := &joinError{
		errs: errs,
	}

	return je
}

// Unwrap returns the inner errors.
//
// Returns:
//   - []error: The inner errors.
func (je joinError) Unwrap() []error {
	if len(je.errs) == 0 {
		return nil
	}

	errs := make([]error, len(je.errs))
	copy(errs, je.errs)

	return errs
}
