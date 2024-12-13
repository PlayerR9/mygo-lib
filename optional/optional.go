package optional

import "github.com/PlayerR9/mygo-lib/common"

// Optional is an interface that represents an optional value.
type Optional interface {
	// IsPresent checks whether the Optional has a value.
	//
	// Returns:
	//   - bool: True if the Optional has a value, false otherwise.
	IsPresent() bool

	// Get returns the value of the Optional if it is present, or an error if
	// the Optional has no value.
	//
	// Returns:
	//   - any: The value of the Optional if present.
	//   - error: An error if the Optional has no value.
	//
	// Errors:
	//   - ErrMissingValue: If the Optional has no value.
	Get() (any, error)
}

// MustGet calls o.Get() and panics if o is nil or if the call to o.Get() returns an error.
//
// Panics:
//   - If o is nil.
//   - If o.Get() returns an error.
//   - If the value returned by o.Get() is not of type E.
//
// Returns:
//   - E: The value returned by o.Get() if the call to o.Get() succeeds.
func MustGet[E any](o Optional) E {
	if o == nil {
		panic(common.NewErrNilParam("o"))
	}

	v, err := o.Get()
	if err != nil {
		panic(ErrMissingValue)
	}

	e, ok := v.(E)
	if !ok {
		panic("value is not of type E")
	}

	return e
}
