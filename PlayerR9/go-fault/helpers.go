package faults

// Innermost returns the innermost base of the given fault.
//
// Parameters:
//   - fault: The fault.
//
// Returns:
//   - Fault: The innermost base of the fault. Nil if the given fault is nil.
func Innermost(fault Fault) Fault {
	if fault == nil {
		return nil
	}

	for {
		inner := fault.Embeds()
		if inner == nil {
			return fault
		}

		fault = inner
	}
}
