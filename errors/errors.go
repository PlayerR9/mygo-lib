package errors

import "fmt"

// ErrPanic occurs when a panic occurs.
type ErrPanic struct {
	// Value is the value of the panic.
	Value any
}

// Error implements error.
func (e ErrPanic) Error() string {
	msg := fmt.Sprintf("panic: %v", e.Value)
	return msg
}

// NewErrPanic returns an error with the given value.
//
// Parameters:
//   - value: The value of the panic.
//
// Returns:
//   - error: An instance of ErrPanic. Never returns nil.
//
// Format:
//
//	"panic: %v"
//
// Where:
//   - %v is the value of the panic.
func NewErrPanic(value any) error {
	e := &ErrPanic{
		Value: value,
	}

	return e
}
