package faults

// Blueprint is a blueprint for a Fault.
type Blueprint interface {
	// String returns the major category of the Fault.
	//
	// Returns:
	//   - string: The major category of the Fault.
	String() string

	// Init creates a new Fault with the given message.
	//
	// Parameters:
	//   - msg: The message for the Fault.
	//
	// Returns:
	//   - Fault: The new Fault. Never returns nil.
	Init(msg string) Fault
}
