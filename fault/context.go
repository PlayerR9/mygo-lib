package fault

import (
	"github.com/PlayerR9/mygo-lib/common"
)

// AddContext adds key-value context to an existing error, converting it into a Fault if necessary.
//
// Parameters:
//   - target: A pointer to the error to which context will be added. Must not be nil.
//   - key: The key for the context to be added.
//   - value: The value for the context to be added.
//
// Returns:
//   - error: Returns an error if the operation failed.
//
// Errors:
//   - common.ErrBadParam: If target is nil.
//   - any other error: Implementation-specific; wrapped in a Fault with message "could not add context".
func AddContext(target *error, key string, value any) error {
	if target == nil {
		return common.NewErrNilParam("target")
	}

	t := *target
	if t == nil {
		return nil
	}

	var ce Fault

	if tmp, ok := t.(Fault); ok {
		ce = tmp
	} else {
		ce = FaultOf("", t)
	}

	err := ce.AddContext(key, value)
	if err == nil {
		*target = ce
		return nil
	}

	fault := FaultOf("could not add context", err)
	return fault
}
