package tables

import (
	"iter"
)

// BoundlessTable is a boundless table. This means that any operation done
// to out-of-bounds cells will not cause any error.
type BoundlessTable[T any] struct {
	// table is the underlying table.
	table [][]T

	// width is the width of the table.
	width uint

	// height is the height of the table.
	height uint
}

// Row returns an iterator over the rows in the table.
//
// Returns:
//   - iter.Seq2[uint, []T]: An iterator over the rows in the table. Never returns nil.
func (t BoundlessTable[T]) Row() iter.Seq2[uint, []T] {
	return func(yield func(uint, []T) bool) {
		for i := uint(0); i < t.height; i++ {
			if !yield(i, t.table[i]) {
				return
			}
		}
	}
}

// Free cleans up the table. Does nothing if the receiver is nil or if
// is already cleaned up.
func (t *BoundlessTable[T]) Free() {
	if t == nil {
		return
	}

	if len(t.table) == 0 {
		t.height = 0
		t.width = 0

		return
	}

	for i := uint(0); i < t.height; i++ {
		clear(t.table[i])
		t.table = nil
	}

	clear(t.table)
	t.table = nil

	t.width = 0
	t.height = 0
}
