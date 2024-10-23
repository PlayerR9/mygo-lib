package common

// ErrAssertFail occurs when an assertion fails.
type ErrAssertFail struct {
	// Inner is the error that occurred.
	Inner error
}

// Error implements the error interface.
func (e ErrAssertFail) Error() string {
	var reason string

	if e.Inner == nil {
		reason = "something went wrong"
	} else {
		reason = e.Inner.Error()
	}

	return "ASSERT FAIL: " + reason
}

// NewErrAssertFail returns a new ErrAssertFail from the given error.
//
// Parameters:
//   - inner: The error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"ASSERT FAIL: <reason>"
//
// Where, <reason> is the error message of the given error. If nil, "something went wrong" is used.
func NewErrAssertFail(inner error) error {
	return &ErrAssertFail{
		Inner: inner,
	}
}

// Unwrap returns the inner error.
//
// Returns:
//   - error: The inner error.
func (e ErrAssertFail) Unwrap() error {
	return e.Inner
}

// Validate panics if the given element is not valid.
//
// Parameters:
//   - elem: The element to validate.
//
// Panics:
//   - ErrAssertFail: If the element is not valid either because it is nil or the Validate method returns an error.
//
// Note: This function is intended to be used for implementing the assert.Validater interface.
func Validate(elem interface{ Validate() error }) {
	if elem == nil {
		panic(NewErrAssertFail(ErrNilReceiver))
	}

	err := elem.Validate()
	if err != nil {
		panic(NewErrAssertFail(err))
	}
}
