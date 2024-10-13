package faults

import "fmt"

// baseFault is a base class for fault implementations.
type baseFault[C FaultCode] struct {
	// code is the fault code.
	code C

	// msg is the error message.
	msg string
}

// Error implements the Fault interface.
func (f baseFault[C]) Error() string {
	return "(" + f.code.String() + ") " + f.msg
}

// Embeds implements the Fault interface.
func (f baseFault[C]) Embeds() Fault {
	return nil
}

// New creates a new Fault with the given fault code and message.
//
// Parameters:
//   - code: The fault code.
//   - msg: The error message.
//
// Returns:
//   - Fault: The new fault. Never returns nil.
func New[C FaultCode](code C, msg string) Fault {
	return &baseFault[C]{
		code: code,
		msg:  msg,
	}
}

// Newf is the same as New, but with a format string.
//
// Parameters:
//   - code: The fault code.
//   - format: The format string.
//   - args: The arguments for the format string.
//
// Returns:
//   - Fault: The new fault. Never returns nil.
func Newf[C FaultCode](code C, format string, args ...any) Fault {
	msg := fmt.Sprintf(format, args...)

	return &baseFault[C]{
		code: code,
		msg:  msg,
	}
}
