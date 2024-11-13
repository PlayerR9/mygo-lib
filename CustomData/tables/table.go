package tables

import "github.com/PlayerR9/mygo-lib/common"

// Table is a boundless table, meaning that the table can be thought of
// as being surrounded by an infinite sea of zero value cells. Thus, out-of-bounds
// reading and writing will not cause any error.
type Table[T any] struct {
	// table is the underlying table.
	table [][]T

	// width is the width of the table.
	width uint

	// height is the height of the table.
	height uint
}

// Free implements common.Typer.
func (t *Table[T]) Free() {
	if t == nil {
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

// NewTable creates a new Table with a width and height.
//
// Parameters:
//   - width: The width of the table.
//   - height: The height of the table.
//
// Returns:
//   - Table[T]: The new Table.
func NewTable[T any](width, height uint) Table[T] {
	table := make([][]T, 0, height)

	for i := uint(0); i < height; i++ {
		table = append(table, make([]T, width, width))
	}

	return Table[T]{
		table:  table,
		width:  width,
		height: height,
	}
}

// Width returns the width of the table.
//
// Returns:
//   - uint: The width of the table.
func (t Table[T]) Width() uint {
	return t.width
}

// Height returns the height of the table.
//
// Returns:
//   - uint: The height of the table.
func (t Table[T]) Height() uint {
	return t.height
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
//   - error: An error of type common.ErrInvalidObject if the table was freed.
func (t Table[T]) CellAt(x, y uint) (T, error) {
	if x >= t.width || y >= t.height {
		return *new(T), nil
	}

	if t.table == nil {
		return *new(T), common.NewErrInvalidObject("CellAt")
	}

	return t.table[y][x], nil
}

// SetCellAt sets the cell at the specified position. Does nothing if the
// position is out of bounds.
//
// Parameters:
//   - value: The value to set.
//   - x: The x position of the cell.
//   - y: The y position of the cell.
//
// Returns:
//   - error: An error if the cell could not be set.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
//   - common.ErrInvalidObject: If the table was freed.
func (t *Table[T]) SetCellAt(value T, x, y uint) error {
	if t == nil {
		return common.ErrNilReceiver
	}

	if x >= t.width || y >= t.height {
		return nil
	}

	if t.table == nil {
		return common.NewErrInvalidObject("SetCellAt")
	}

	t.table[y][x] = value

	return nil
}

// ResizeHeight resizes the height of the table.
//
// Parameters:
//   - new_height: The new height of the table.
//
// Returns:
//   - error: An error if the table could not be resized.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
//   - common.ErrInvalidObject: If the table was freed.
func (t *Table[T]) ResizeHeight(new_height uint) error {
	if t == nil {
		return common.ErrNilReceiver
	}

	if new_height == t.height {
		return nil
	}

	if t.table == nil {
		return common.NewErrInvalidObject("ResizeHeight")
	}

	if new_height < t.height {
		clear(t.table[t.height:])

		t.table = t.table[:new_height]
	} else {
		for i := t.height; i < new_height; i++ {
			t.table = append(t.table, make([]T, t.width))
		}
	}

	return nil
}

// ResizeWidth resizes the width of the table.
//
// Parameters:
//   - new_width: The new width of the table.
//
// Returns:
//   - error: An error if the table could not be resized.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
//   - common.ErrInvalidObject: If the table was freed.
func (t *Table[T]) ResizeWidth(new_width uint) error {
	if t == nil {
		return common.ErrNilReceiver
	}

	if new_width == t.width {
		return nil
	}

	if t.table == nil {
		return common.NewErrInvalidObject("ResizeWidth")
	}

	if new_width < t.width {
		for i := uint(0); i < t.height; i++ {
			clear(t.table[i][t.width:])
			t.table[i] = t.table[i][:new_width]
		}
	} else {
		for i := uint(0); i < t.height; i++ {
			t.table[i] = append(t.table[i], make([]T, new_width-t.width)...)
		}
	}

	return nil
}
