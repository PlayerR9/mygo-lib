package errors

import "errors"

var (
	// ErrNilReceiver occurs when a method is called on a receiver that was not expected to be
	// nil. This error can be checked with the == operator.
	//
	// Format:
	//
	//		"receiver must not be nil"
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}

// ErrBadParam occurs when a function is called with a parameter that does not
// satisfy the function's contract.
type ErrBadParam struct {
	// ParamName is the name of the parameter that caused the error.
	ParamName string

	// Message is the error message.
	Message string
}

// Error implements error.
func (e ErrBadParam) Error() string {
	var msg string

	if e.Message == "" {
		msg = "is not valid"
	} else {
		msg = e.Message
	}

	if e.ParamName == "" {
		return "parameter " + msg
	} else {
		return "parameter (" + e.ParamName + ") " + msg
	}
}

// NewErrBadParam creates a new ErrBadParam error with the given parameter name and
// message.
//
// Parameters:
//   - param_name: The name of the parameter that caused the error.
//   - msg: The error message.
//
// Returns:
//   - error: The new ErrBadParam error. Never returns nil.
//
// Format:
//
//	"parameter (<param_name>) <msg>"
//
// Where:
//   - <param_name> is the name of the parameter that caused the error.
//   - <msg> is the error message. If empty, the message defaults to "is not valid".
//
// However, if param_name is empty, the format is:
//
//	"parameter <msg>"
//
// Where, <msg> is the error message. If empty, the message defaults to "is not valid".
func NewErrBadParam(param_name, msg string) error {
	err := &ErrBadParam{
		ParamName: param_name,
		Message:   msg,
	}

	return err
}

// NewErrNilParam is a convenience function for creating a new ErrBadParam error
// with the message "must not be nil".
//
// Parameters:
//   - param_name: The name of the parameter that caused the error.
//
// Returns:
//   - error: The new ErrBadParam error. Never returns nil.
//
// Format:
//
//	"parameter (<param_name>) must not be nil"
//
// Where, <param_name> is the name of the parameter that caused the error.
//
// However, if param_name is empty, the format is:
//
//	"parameter must not be nil"
func NewErrNilParam(param_name string) error {
	err := &ErrBadParam{
		ParamName: param_name,
		Message:   "must not be nil",
	}

	return err
}
