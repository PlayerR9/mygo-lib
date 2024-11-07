package listlike

import (
	"errors"
)

var (
	// ErrEmptyStack occurs when a pop or peek operation is called on an empty stack.
	// This error can be checked with the == operator.
	//
	// Format:
	// 	"empty stack"
	ErrEmptyStack error

	// ErrFullStack occurs when a push operation is called on a full stack.
	// This error can be checked with the == operator.
	//
	// Format:
	// 	"full stack"
	ErrFullStack error

	// ErrEmptyQueue occurs when a pop or peek operation is called on an empty queue.
	// This error can be checked with the == operator.
	//
	// Format:
	// 	"empty queue"
	ErrEmptyQueue error

	// ErrCannotPush occurs when a push operation is called on a refusable stack that
	// was not accepted nor refused yet. This error can be checked with the == operator.
	//
	// Format:
	// 	"cannot push elements: stack not accepted nor refused"
	ErrCannotPush error
)

func init() {
	ErrEmptyStack = errors.New("empty stack")
	ErrFullStack = errors.New("full stack")
	ErrEmptyQueue = errors.New("empty queue")
	ErrCannotPush = errors.New("cannot push elements: stack not accepted nor refused")
}

type ErrBadImplement struct {
	MethodName string
	Inner      error
}

func (e ErrBadImplement) Error() string {
	var msg string

	if e.Inner != nil {
		msg = ": " + e.Inner.Error()
	}

	return "method " + e.MethodName + " is badly implemented" + msg
}

func NewErrBadImplement(method_name string, inner error) error {
	return &ErrBadImplement{
		MethodName: method_name,
		Inner:      inner,
	}
}
