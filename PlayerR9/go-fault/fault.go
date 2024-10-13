package faults

// FaultCode is the type of a fault code.
type FaultCode interface {
	~int

	// String returns the string representation of the fault code.
	//
	// Returns:
	//   - string: The string representation of the fault code.
	String() string
}

// Fault is a fault.
type Fault interface {
	// Embeds returns the base of the fault.
	//
	// Returns:
	//   - Fault: The base of the fault.
	Embeds() Fault

	// Error returns the error message of the fault.
	//
	// Returns:
	//   - string: The error message of the fault.
	Error() string
}

// Is checks whether the given fault is of the same type as the target fault.
//
// Two faults are said to be equal if they have the same pointer value or if the fault
// implements the IsFault(Fault) bool method such that fault.IsFault(target) returns true. In
// any other case, IsFault returns false.
//
// Parameters:
//   - fault: The fault to check.
//   - target: The target fault.
//
// Returns:
//   - bool: True if the given fault is of the same type as the target fault, false otherwise.
func Is(fault Fault, target Fault) bool {
	if fault == nil || target == nil {
		return false
	} else if fault == target {
		return true
	}

	flt, ok := fault.(interface{ IsFault(Fault) bool })
	if !ok {
		return false
	}

	return flt.IsFault(target)
}
