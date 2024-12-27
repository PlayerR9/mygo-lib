package errors

import (
	"reflect"
	"strings"

	"github.com/PlayerR9/mygo-lib/errors/internal"
)

const (
	// DefaultErrorMessage is the default error message.
	DefaultErrorMessage string = "something went wrong"
)

// New returns a new error with the given message. If the message is empty, it
// defaults to DefaultErrorMessage.
//
// Parameters:
//   - msg: The error message.
//
// Returns:
//   - error: The new error. Never returns nil.
func New(msg string) error {
	if msg == "" {
		msg = DefaultErrorMessage
	}

	be := internal.NewBaseError(msg)

	return be
}

// Errorf returns an error with a formatted message.
//
// Parameters:
//   - format: The format string.
//   - args: The arguments for the format string.
//
// Returns:
//   - error: The new error instance. Never returns nil.
func Errorf(format string, args ...any) error {
	format = strings.ReplaceAll(format, "%w", "%s")

	fe := internal.NewFormatError(format, args)
	return fe
}

// Reverse reverses a slice of errors in-place.
//
// Parameters:
//   - errs: The slice of errors to be reversed.
func Reverse(errs []error) {
	if len(errs) == 0 {
		return
	}

	internal.Reverse(errs)
}

// Unwrap extracts inner errors from the provided error.
//
// It checks if the error implements either the `Unwrap() error` or
// `Unwrap() []error` interface and returns a slice of errors.
//
// Parameters:
//   - target: The error from which to extract inner errors.
//
// Returns:
//   - []error: A slice of unwrapped errors. Returns nil if the target is nil or
//     if no errors are unwrapped.
func Unwrap(target error) []error {
	if target == nil {
		return nil
	}

	errs := internal.Unwrap(target)
	return errs
}

// Is reports whether any error in err's chain matches target.
//
// Parameters:
//   - err: The error to be checked. Can be nil.
//   - target: The target error to match against. Must not be nil.
//
// Returns:
//   - bool: True if any error in err's chain matches target, false otherwise.
func Is(err error, target error) bool {
	if err == nil || target == nil {
		return false
	}

	stack := []error{err}

	isComparable := reflect.TypeOf(target).Comparable()
	var found bool

	if isComparable {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top == target {
				found = true
				break
			}

			x, ok := top.(interface{ Is(error) bool })
			ok = ok && x.Is(target)
			if ok {
				found = true
				break
			}

			errs := internal.Unwrap(top)
			if len(errs) == 0 {
				continue
			}

			internal.Reverse(errs)
			stack = append(stack, errs...)
		}
	} else {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			x, ok := top.(interface{ Is(error) bool })
			ok = ok && x.Is(target)
			if ok {
				return true
			}

			errs := internal.Unwrap(top)
			if len(errs) == 0 {
				continue
			}

			internal.Reverse(errs)
			stack = append(stack, errs...)
		}
	}

	return found
}

// CountNonNilErrors counts the number of non-nil errors in the given slice.
//
// Parameters:
//   - errs: The slice of errors to count.
//
// Returns:
//   - uint: The number of non-nil errors in the slice.
func CountNonNilErrors(errs []error) uint {
	if len(errs) == 0 {
		return 0
	}

	n := internal.Count(errs)
	return n
}

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

	count := internal.Count(errs)
	if count == 0 {
		return nil
	}

	result := make([]error, count)
	copy(result, errs)

	internal.RejectNilErrors(&result, count)
	return result
}

// Join joins a slice of errors into one error. If the slice is empty, it
// returns nil. If the slice contains any nil errors, they are ignored.
// If the slice contains only one error, it is returned as is. Otherwise, it
// returns an error that contains all the non-nil errors in the slice.
//
// It uses the internal/internal.JoinError type to represent the error. This
// type implements the error interface and contains a slice of errors.
//
// Parameters:
//   - errs: A slice of error interfaces, which may include nil values.
//
// Returns:
//   - error: An error that contains all the non-nil errors in the slice. If the
//     slice is empty or only contains nil errors, it returns nil.
func Join(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	count := internal.Count(errs)
	if count == 0 {
		return nil
	}

	result := make([]error, count)
	copy(result, errs)

	internal.RejectNilErrors(&result, count)

	if count == 1 {
		return result[0]
	}

	je := internal.NewJoinError(errs)
	return je
}
