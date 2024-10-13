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
// If the target fault is nil, Is returns false.
//
// Otherwise, Is returns true if the given fault is of the same type as the target fault,
// i.e. if the given fault's error code is the same as the fault code of the target fault.
//
// If the given fault does not have an error code, Is checks whether the given fault
// is the same as the target fault or whether the given fault has the target fault as its base.
//
// If the given fault has an error code and the target fault does not, Is returns false.
//
// Parameters:
//   - fault: The fault to check.
//   - target: The target fault.
//
// Returns:
//   - bool: True if the given fault is of the same type as the target fault, false otherwise.
func Is(fault Fault, target Fault) bool {
	if target == nil {
		return false
	}

	for fault != nil {
		if fault == target {
			return true
		}

		flt, ok := fault.(interface{ IsFault(Fault) bool })
		if ok {
			return flt.IsFault(target)
		}

		fault = fault.Embeds()
	}

	return false
}
