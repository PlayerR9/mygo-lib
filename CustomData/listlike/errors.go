package listlike

import "errors"

var (
	// ErrEmptyStack occurs when a pop or peek operation is called on an empty stack.
	// This error can be checked with the == operator.
	//
	// Format:
	// 	"empty stack"
	ErrEmptyStack error

	// ErrEmptyQueue occurs when a pop or peek operation is called on an empty queue.
	// This error can be checked with the == operator.
	//
	// Format:
	// 	"empty queue"
	ErrEmptyQueue error
)

func init() {
	ErrEmptyStack = errors.New("empty stack")
	ErrEmptyQueue = errors.New("empty queue")
}
