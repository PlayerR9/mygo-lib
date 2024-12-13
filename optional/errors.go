package optional

import (
	flt "github.com/PlayerR9/mygo-lib/go-fault"
)

var (
	// bltErrMissingValue is the blueprint for ErrMissingValue.
	bltErrMissingValue flt.Blueprint
)

func init() {
	bltErrMissingValue = flt.New("")
}

// ErrMissingValue returns an error that occurs when an optional value is missing.
//
// Format:
//
//	"missing value"
func ErrMissingValue() flt.Fault {
	return bltErrMissingValue.Init("missing value")
}
