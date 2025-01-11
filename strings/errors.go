package strings

import "errors"

var (
	// ErrNoPredicate occurs when a predicate is not provided.
	//
	// Format:
	// 	"no predicate was provided"
	ErrNoPredicate error
)

func init() {
	ErrNoPredicate = errors.New("no predicate was provided")
}
