package list

import "errors"

var (
	// ErrEmptyList occurs when a delist or front operation is called on an empty list.
	// This error can be checked with the == operator.
	ErrEmptyList error

	// ErrFullList occurs when a enlist operation is called on a full list.
	// This error can be checked with the == operator.
	ErrFullList error
)

func init() {
	ErrEmptyList = errors.New("empty list")
	ErrFullList = errors.New("full list")
}
