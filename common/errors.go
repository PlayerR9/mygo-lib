package common

import "errors"

var (
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}

type ErrBadParameter struct {
	ParamName string
	Msg       string
}

func (e ErrBadParameter) Error() string {
	var msg string

	if e.Msg == "" {
		msg = "is invalid"
	} else {
		msg = e.Msg
	}

	if e.ParamName == "" {
		return "Parameter " + msg
	} else {
		return "Parameter (" + e.ParamName + ") " + msg
	}
}

func NewErrBadParameter(param_name, msg string) error {
	return &ErrBadParameter{
		ParamName: param_name,
		Msg:       msg,
	}
}

func NewErrNilParameter(param_name string) error {
	return &ErrBadParameter{
		ParamName: param_name,
		Msg:       "must not be nil",
	}
}
