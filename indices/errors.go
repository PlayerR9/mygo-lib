package indices

import "errors"

var (
	// ErrMissingValue occurs when a value is not present.
	//
	// Format:
	// 	"value is not present"
	ErrMissingValue error
)

func init() {
	ErrMissingValue = errors.New("value is not present")
}
