package fault

// Fault is an interface that represents an error.
type Fault interface {
	// Error returns the error message.
	//
	// Returns:
	//   - string: The error message.
	Error() string

	// Init creates a new Fault with the given message and who shares the
	// same blueprint.
	//
	// Parameters:
	//   - msg: The message for the Fault.
	//
	// Returns:
	//   - Fault: The new Fault. Never returns nil.
	Init(msg string) Fault
}

// NewFault creates a new Fault with the given name and message.
//
// Parameters:
//   - name: The name of the Fault.
//   - msg: The message for the Fault.
//
// Returns:
//   - Fault: The new Fault. Never returns nil.
func NewFault(name, msg string) Fault {
	blueprint := New(name)
	fault := blueprint.Init(msg)
	return fault
}
