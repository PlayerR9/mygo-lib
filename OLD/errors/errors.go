package errors

import (
	"fmt"
)

// ErrPanic is an error that is returned when a panic occurs.
type ErrPanic struct {
	// value is the value that was passed to panic.
	value any
}

// Error implements the error interface.
func (e ErrPanic) Error() string {
	return fmt.Sprintf("panic: %v", e.value)
}

// NewErrPanic creates a new ErrPanic error.
//
// Parameters:
//   - value: The value that was passed to panic.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"panic: <value>"
//
// Where:
//   - <value>: The value that was passed to panic.
func NewErrPanic(value any) error {
	return &ErrPanic{
		value: value,
	}
}
