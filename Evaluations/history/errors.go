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

	// ErrErrorEncountered is an error that is returned when an error is encountered but
	// it is not critical enough to stop the evaluation.
	ErrErrorEncountered error

	// ErrExhausted is an error that is returned when the iterator is exhausted.
	ErrExhausted error
)

func init() {
	ErrHistoryDone = errors.New("history is done")
	ErrErrorEncountered = errors.New("error encountered")
	ErrExhausted = errors.New("iterator is exhausted")
}
