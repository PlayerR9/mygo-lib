package optional

import "errors"

var (
	// ErrMissingValue is an error that occurs when an optional value is missing.
	//
	// Format:
	//
	//	"missing value"
	ErrMissingValue error
)

func init() {
	ErrMissingValue = errors.New("missing value")
}
