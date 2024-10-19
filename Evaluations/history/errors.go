package history

import (
	"errors"
)

var (
	// ErrHistoryDone is an error that is returned when the history is done.
	//
	// Format:
	//   "history is done"
	ErrHistoryDone error
)

func init() {
	ErrHistoryDone = errors.New("history is done")
}
