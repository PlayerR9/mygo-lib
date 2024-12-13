package common

import (
	flt "github.com/PlayerR9/mygo-lib/go-fault"
)

var (
	// bltErrBadParam is the blueprint for a bad parameter error.
	bltErrBadParam flt.Blueprint

	// bltErrOperationFailed is the blueprint for an operation failed error.
	bltErrOperationFailed flt.Blueprint
)

func init() {
	bltErrBadParam = flt.New("Bad Parameter")
	bltErrOperationFailed = flt.New("Operation Failed")
}

// ErrNilReceiver returns a fault that occurs when a method is called on a receiver that was expected
// to be non-nil.
//
// Returns:
//   - flt.Fault: An instance of a fault indicating that the receiver must not be nil. Never returns nil.
//
// Format:
//
//	"receiver must not be nil"
func ErrNilReceiver() flt.Fault {
	fault := bltErrOperationFailed.Init("receiver must not be nil")
	return fault
}

// ErrBadParam returns an error that occurs when a parameter is bad. (i.e., not a valid value).
//
// Parameters:
//   - param_name: The name of the parameter causing the error.
//   - msg: The error message describing why the parameter is bad. If empty, "is not valid" is used.
//
// Returns:
//   - error: An instance of a bad parameter error. Never returns nil.
//
// Format:
//
//	"parameter (<param_name>) <msg>"
//
// where:
//   - (<param_name>): The name of the parameter. If empty, it is omitted.
//   - <msg>: The error message describing why the parameter is bad. If empty, "is not valid" is used.
func ErrBadParam(param_name, msg string) flt.Fault {
	if msg == "" {
		msg = "is not valid"
	}

	var fault flt.Fault

	if param_name == "" {
		fault = bltErrBadParam.Init("parameter " + msg)
	} else {
		fault = bltErrBadParam.Init("parameter (" + param_name + ") " + msg)
	}

	return fault
}

// ErrBadParam returns an error that occurs when a parameter is bad. (i.e., not a valid value).
//
// Parameters:
//   - param_name: The name of the parameter causing the error.
//
// Returns:
//   - error: An instance of a bad parameter error. Never returns nil.
//
// Format:
//
//	"parameter (<param_name>) must not be nil"
//
// where:
//   - (<param_name>): The name of the parameter. If empty, it is omitted.
func ErrNilParam(param_name string) flt.Fault {
	var fault flt.Fault

	if param_name == "" {
		fault = bltErrBadParam.Init("parameter must not be nil")
	} else {
		fault = bltErrBadParam.Init("parameter (" + param_name + ") must not be nil")
	}

	return fault
}
