package indices

import (
	gers "github.com/PlayerR9/mygo-lib/errors"
)

var (
	// ErrMissingValue occurs when a value is not present.
	//
	// Format:
	// 	"value is not present"
	ErrMissingValue error
)

func init() {
	ErrMissingValue = gers.New("value is not present")
}
