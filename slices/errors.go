package slices

import "errors"

var (
	// ErrNilReceiver occurs when a method is called on a receiver that was not expected to be
	// nil. This error can be checked with the == operator.
	//
	// Format:
	//		"receiver must not be nil"
	ErrNilReceiver error

	// ErrNoPredicate occurs when a predicate is not provided.
	//
	// Format:
	// 	"no predicate was provided"
	ErrNoPredicate error
)

func init() {
	ErrNoPredicate = errors.New("no predicate was provided")
}
