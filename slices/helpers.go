package slices

import flt "github.com/PlayerR9/mygo-lib/fault"

// MayGetter is a getter that may return an error.
type MayGetter[E any] interface {
	// Get returns the value or an error.
	//
	// Returns:
	//   - E: The value or its zero value if the value is not present.
	//   - fault.Fault: An error if the value is not present.
	//
	// Errors:
	//   - ErrMissingValue: If the value is not present.
	Get() (E, flt.Fault)
}

// MustGet calls mg.Get() and panics if mg is nil or if the call to mg.Get() returns an error.
//
// Panics:
//   - If mg is nil.
//   - If mg.Get() returns an error.
//
// Returns:
//   - E: The value returned by mg.Get().
func MustGet[E any](mg MayGetter[E]) E {
	if mg == nil {
		panic("parameter (mg) must not be nil")
	}

	v, err := mg.Get()
	if err != nil {
		panic(err)
	}

	return v
}
