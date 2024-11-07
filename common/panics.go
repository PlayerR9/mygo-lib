package common

import "fmt"

type ErrMethodNotImpl struct {
	MethodName string
	Receiver   any
}

func (e ErrMethodNotImpl) Error() string {
	var msg string

	if e.Receiver != nil {
		msg = fmt.Sprintf(" by %T", e.Receiver)
	}

	return "method " + e.MethodName + " is not implemented" + msg
}

func NewErrMethodNotImpl(method_name string, receiver any) error {
	return &ErrMethodNotImpl{
		MethodName: method_name,
		Receiver:   receiver,
	}
}
