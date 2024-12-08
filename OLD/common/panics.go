package common

import "fmt"

// ErrMethodNotImpl occurs when a method is not implemented.
type ErrMethodNotImpl struct {
	// MethodName is the name of the method that is not implemented.
	MethodName string

	// Receiver is the receiver of the method that is not implemented.
	Receiver any
}

// Error implements the error interface.
func (e ErrMethodNotImpl) Error() string {
	var msg string

	if e.Receiver != nil {
		msg = fmt.Sprintf(" by %T", e.Receiver)
	}

	return "method " + e.MethodName + " is not implemented" + msg
}

// NewErrMethodNotImpl creates a new ErrMethodNotImpl error with the specified method name and receiver.
//
// Parameters:
//   - method_name: The name of the method that is not implemented.
//   - receiver: The receiver of the method that is not implemented.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"method <method_name> is not implemented by <receiver>"
//
// where:
//   - <method_name>: The name of the method that is not implemented.
//   - <receiver>: The receiver of the method that is not implemented. If nil, it is omitted.
func NewErrMethodNotImpl(method_name string, receiver any) error {
	return &ErrMethodNotImpl{
		MethodName: method_name,
		Receiver:   receiver,
	}
}
