package internal

import "errors"

var (
	// BugMismatchWritten occurs when n != len(data) when err == nil.
	//
	// Format:
	// 	"[bug]: n != len(data) when err == nil"
	BugMismatchWritten error
)

func init() {
	BugMismatchWritten = errors.New("[bug]: n != len(data) when err == nil")
}
