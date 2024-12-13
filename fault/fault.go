package fault

import "github.com/PlayerR9/mygo-lib/common"

// Fault is the interface for a fault, that is, an error with context.
type Fault interface {
	// AddContext adds context to the fault.
	//
	// Parameters:
	//   - key: The key to add to the context.
	//   - value: The value to add to the context.
	//
	// Returns:
	//   - error: A fault if the operation failed.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error: Implementation-specific.
	AddContext(key string, value any) error

	error
}

// baseFault is a type of error that contains information about the state
// of the program.
type baseFault struct {
	// msg is a top-level error message.
	msg string

	// inner is the error that the fault was caused by.
	inner error

	// context is the context of the error. This is initialized
	// at the first call to AddContext.
	context map[string]any
}

// AddContext implements Fault.
func (f *baseFault) AddContext(key string, value any) error {
	if f == nil {
		return common.ErrNilReceiver
	}

	if f.context == nil {
		f.context = make(map[string]any)
	}

	f.context[key] = value

	return nil
}

// Error implements error.
func (f baseFault) Error() string {
	return f.msg
}

// FaultOf creates a new fault instance from the given target error.
//
// Parameters:
//   - msg: The message for the fault.
//   - target: The original error to be wrapped in a fault.
//
// Returns:
//   - Fault: A new fault instance wrapping the target error. Never returns nil.
//
// Behaviors:
//   - If the message is empty, it uses the default error message.
func FaultOf(msg string, target error) Fault {
	if msg == "" {
		msg = DefaultErrorMessage
	}

	f := &baseFault{
		msg:   msg,
		inner: target,
	}

	return f
}
