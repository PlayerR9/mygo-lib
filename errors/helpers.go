package errors

// RejectNilErrors removes all nil errors from the given slice.
//
// It iterates over the provided slice, counting the number of non-nil errors.
// If the slice is empty or only contains nil errors, it returns nil.
// Otherwise, it creates a new slice containing the non-nil errors and returns it.
//
// Parameters:
//   - errs: A slice of error interfaces, which may include nil values.
//
// Returns:
//   - []error: A new slice containing only the non-nil errors from the original slice.
func RejectNilErrors(errs []error) []error {
	if len(errs) == 0 {
		return nil
	}

	var count uint

	for _, err := range errs {
		if err != nil {
			count++
		}
	}

	if count == 0 {
		return nil
	}

	result := make([]error, 0, count)

	for _, err := range errs {
		if err != nil {
			result = append(result, err)
		}
	}

	return result
}
