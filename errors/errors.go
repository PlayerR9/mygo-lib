package errors

// ErrWhile is an error that occurs while doing something.
type ErrWhile struct {
	// Process is the process in which the error occurred.
	Process string

	// Inner is the original error.
	Inner error
}

// Error implements error.
func (e ErrWhile) Error() string {
	var process string

	if e.Process == "" {
		process = "doing something"
	} else {
		process = e.Process
	}

	var msg string

	if e.Inner == nil {
		msg = "something went wrong while " + process + "..."
	} else {
		reason := e.Inner.Error()
		msg = "while " + process + ": " + reason
	}

	return msg
}

// NewErrWhile returns an error that occurs while doing something.
//
// Parameters:
//   - process: The process in which the error occurred.
//   - inner: The original error.
//
// Returns:
//   - error: An instance of ErrWhile. Never returns nil.
//
// Format:
//
//	"while <process>: <inner>"
//
// Where:
//   - <process> is the process in which the error occurred.
//   - <inner> is the original error message. If nil, "something went wrong" is used.
func NewErrWhile(process string, inner error) error {
	e := &ErrWhile{
		Process: process,
		Inner:   inner,
	}

	return e
}

// Unwrap returns the inner error.
//
// Returns:
//   - error: The inner error instance.
func (e ErrWhile) Unwrap() error {
	return e.Inner
}
