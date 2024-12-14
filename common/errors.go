package common

import "errors"

var (
	// ErrNilReceiver occurs when a method is called on a receiver that is nil. This
	// error can be checked with the == operator.
	//
	// Format:
	// 	"receiver must not be nil"
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}

// ErrBadParam occurs when a parameter is not valid.
type ErrBadParam struct {
	// Name is the name of the parameter.
	Name string

	// Msg is the error message.
	Msg string
}

// Error implements error.
func (e ErrBadParam) Error() string {
	var msg string

	if e.Msg == "" {
		msg = "is not valid"
	} else {
		msg = e.Msg
	}

	if e.Name == "" {
		return "parameter " + msg
	} else {
		return "parameter (" + e.Name + ") " + msg
	}
}

// NewErrBadParam creates and returns a new ErrBadParam error with the specified
// parameter name and error message.
//
// Parameters:
//   - name: The name of the parameter.
//   - msg: The error message.
//
// Returns:
//   - error: A pointer to the newly created ErrBadParam. Never returns nil.
//
// Format:
//
//	"parameter (<name>) <msg>"
//
// Where:
//   - <name> is the name of the parameter. If empty, it is omitted.
//   - <msg> is the error message. If empty, defaults to "is not valid".
func NewErrBadParam(name, msg string) error {
	err := &ErrBadParam{
		Name: name,
		Msg:  msg,
	}
	return err
}

// NewErrNilParam creates and returns a new ErrBadParam error with the specified
// parameter name and error message "must not be nil".
//
// Parameters:
//   - name: The name of the parameter.
//
// Returns:
//   - error: A pointer to the newly created ErrBadParam. Never returns nil.
//
// Format:
//
//	"parameter (<name>) must not be nil"
//
// Where, <name> is the name of the parameter. If empty, it is omitted.
func NewErrNilParam(name string) error {
	err := &ErrBadParam{
		Name: name,
		Msg:  "must not be nil",
	}
	return err
}
