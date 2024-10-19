package history

import (
	"errors"
)

var (
	ErrHistoryDone error
)

func init() {
	ErrHistoryDone = errors.New("history is done")
}
