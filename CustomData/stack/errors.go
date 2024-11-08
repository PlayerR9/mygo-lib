package stack

import "errors"

var (
	// ErrEmptyStack occurs when a pop or peek operation is called on an empty stack.
	// This error can be checked with the == operator.
	ErrEmptyStack error

	// ErrFullStack occurs when a push operation is called on a full stack.
	// This error can be checked with the == operator.
	ErrFullStack error

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
	ErrCannotPush = errors.New("cannot push elements: stack not accepted nor refused")
}
