package tables

import (
	"iter"

	"github.com/PlayerR9/mygo-lib/common"
)

// Table is a boundless table. This means that any operation done
// to out-of-bounds cells will not cause any error.
type Table[T any] struct {
	// table is the underlying table.
	table [][]T

	// width is the width of the table.
	width uint

	// height is the height of the table.
	height uint
}

// NewTable creates a new Table with a width and height.
//
// Parameters:
//   - width: The width of the table.
//   - height: The height of the table.
//
// Returns:
//   - *Table: The new Table. Never returns nil.
func NewTable[T any](width, height uint) *Table[T] {
	table := make([][]T, 0, height)

	for i := uint(0); i < height; i++ {
		table = append(table, make([]T, width, width))
	}

	return &Table[T]{
		table:  table,
		width:  width,
		height: height,
	}
}

// Height returns the height of the table.
//
// Returns:
//   - uint: The height of the table.
func (t Table[T]) Height() uint {
	return t.height
}

// Width returns the width of the table.
//
// Returns:
//   - uint: The width of the table.
func (t Table[T]) Width() uint {
	return t.width
}

// CellAt returns the cell at the specified position.
//
// Parameters:
//   - x: The x position of the cell.
//   - y: The y position of the cell.
//
// Returns:
//   - T: The cell at the specified position. The zero value if the position
//     is out of bounds.
func (t Table[T]) CellAt(x, y uint) T {
	if x >= t.width || y >= t.height {
		return *new(T)
	}

	return t.table[y][x]
}

// ResizeWidth resizes the width of the table. The width is not
// resized if the receiver is nil or the new width is the same as the
// current width.
//
// Parameters:
//   - new_width: The new width of the table.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (t *Table[T]) ResizeWidth(new_width uint) error {
	if t == nil {
		return common.ErrNilReceiver
	}

	if new_width == t.width {
		return nil
	}

	if new_width < t.width {
		for i := uint(0); i < t.height; i++ {
			t.table[i] = t.table[i][:new_width:new_width]
		}
	} else {
		extension := make([]T, new_width-t.width)

		for i := uint(0); i < t.height; i++ {
			t.table[i] = append(t.table[i], extension...)
		}
	}

	return nil
}

// ResizeHeight resizes the height of the table. The height is not
// resized if the receiver is nil or the new height is the same as the
// current height.
//
// Parameters:
//   - new_height: The new height of the table.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (t *Table[T]) ResizeHeight(new_height uint) error {
	if t == nil {
		return common.ErrNilReceiver
	}

	if new_height == t.height {
		return nil
	}

	if new_height < t.height {
		t.table = t.table[:new_height:new_height]
	} else {
		for i := t.height; i < new_height; i++ {
			t.table = append(t.table, make([]T, t.width))
		}
	}

	return nil
}

// SetCellAt sets the cell at the specified position. The cell is not
// set if the receiver is nil or the position is out of bounds.
//
// Parameters:
//   - cell: The cell to set.
//   - x: The x position of the cell.
//   - y: The y position of the cell.
func (t *Table[T]) SetCellAt(cell T, x, y uint) {
	if t == nil || y >= t.height || x >= t.width {
		return
	}

	t.table[y][x] = cell
}

// Row returns an iterator over the rows in the table.
//
// Returns:
//   - iter.Seq2[uint, []T]: An iterator over the rows in the table. Never returns nil.
func (t Table[T]) Row() iter.Seq2[uint, []T] {
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
func (t *Table[T]) Free() {
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
