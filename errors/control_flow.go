package errors

import "github.com/PlayerR9/mygo-lib/common"

// Throw panics with the given error, if it is not nil.
//
// Parameters:
//   - err: The error to panic with.
func Throw(err error) {
	if err == nil {
		return
	}

	panic(err)
}

// makeCatchFn returns a function that catches any panic and assigns the error to
// the address provided by the `caught` parameter. The returned function is
// intended to be used as a deferred function.
//
// The assignment is done only if the recovered value is not nil. If the recovered
// value is a string, an error is created with that string. If the recovered value
// is an error, it is assigned directly to the `caught` address. If the recovered
// value is neither a string nor an error, an ErrPanic is created with the
// recovered value and assigned to the `caught` address.
//
// Parameters:
//   - caught: A pointer to an error to assign the caught error to.
//
// Returns:
//   - func(): A function that catches any panic and assigns the error to the
//     address provided by the `caught` parameter. Never returns nil.
func makeCatchFn(caught *error) func() {
	fn := func() {
		r := recover()
		if r == nil {
			return
		}

		var err error

		switch r := r.(type) {
		case string:
			err = common.New(r)
		case error:
			err = r
		default:
			err = NewErrPanic(r)
		}

		*caught = err
	}

	return fn
}

// Try calls the provided function and returns any error that occurs if the function
// panics. If the function does not panic, the function returns nil.
//
// Parameters:
//   - panic_fn: A function that may panic.
//
// Returns:
//   - error: An error if the function panics, otherwise nil.
//
// Errors:
//   - any error that occurs if the function panics.
//
// If the panic value is a string, an error is created with that string. If the
// panic value is an error, it is assigned directly to the `caught` address. If
// the panic value is neither a string nor an error, an ErrPanic is created with
// the panic value and assigned to the `caught` address.
func Try(panic_fn func()) error {
	if panic_fn == nil {
		return nil
	}

	var caught error

	catchFn := makeCatchFn(&caught)

	fn := func() {
		defer catchFn()
		panic_fn()
	}

	fn()

	return caught
}
