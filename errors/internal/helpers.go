package internal

// Reverse reverses a slice of errors in-place.
//
// Parameters:
//   - errs: The slice of errors to be reversed.
func Reverse(errs []error) {
	j := len(errs) - 1

	for i := 0; i < j; i++ {
		errs[i], errs[j] = errs[j], errs[i]
		j--
	}
}

// Count counts the number of non-nil errors in the given slice.
//
// Parameters:
//   - errs: The slice of errors to count.
//
// Returns:
//   - uint: The number of non-nil errors in the slice.
func Count(errs []error) uint {
	var count uint

	for _, err := range errs {
		if err != nil {
			count++
		}
	}

	return count
}

// RejectNilErrors rejects in-place all nil errors from the given slice and returns
// a new slice containing only the non-nil errors. The length of the new slice is
// given by the count parameter.
//
// Parameters:
//   - errs: The slice of errors to filter.
//   - count: The length of the new slice.
func RejectNilErrors(errs *[]error, count uint) {
	var end uint

	for i := 0; i < len(*errs) && end < count; i++ {
		err := (*errs)[i]
		if err == nil {
			continue
		}

		(*errs)[end] = err
		end++
	}

	clear((*errs)[end:])
	*errs = (*errs)[:end]
}

// Unwrap returns a slice of errors extracted from the target error.
//
// It checks if the target error implements either the `Unwrap() error` or `Unwrap() []error`
// interface. If the target implements `Unwrap() error`, it returns a slice containing the
// single unwrapped error. If the target implements `Unwrap() []error`, it returns the slice
// of unwrapped errors, filtering out any nil errors.
//
// Parameters:
//   - target: The error from which to extract inner errors.
//
// Returns:
//   - []error: A slice of unwrapped errors. Returns nil if no errors are unwrapped or if the
//     target does not implement any `Unwrap` interface.
func Unwrap(target error) []error {
	var result []error

	switch target := target.(type) {
	case interface{ Unwrap() error }:
		err := target.Unwrap()
		if err != nil {
			result = make([]error, 0, 1)
			result = append(result, err)
		}
	case interface{ Unwrap() []error }:
		result = target.Unwrap()

		count := Count(result)
		if count == 0 {
			return nil
		}

		RejectNilErrors(&result, count)
	}

	return result
}
