package optional

import flt "github.com/PlayerR9/mygo-lib/go-fault"

// someOptional is an Optional that has a value.
type someOptional struct {
	// value is the value of the Optional.
	value any
}

// IsPresent implements Optional.
//
// Always returns true.
func (so someOptional) IsPresent() bool {
	return true
}

// Get implements Optional.
//
// Never returns an error.
func (so someOptional) Get() (any, flt.Fault) {
	return so.value, nil
}

// Some creates an Optional with the given value.
//
// Parameters:
//   - value: The value to be wrapped in an Optional.
//
// Returns:
//   - Optional: An Optional containing the provided value. Never returns nil.
func Some(value any) Optional {
	return someOptional{
		value: value,
	}
}
