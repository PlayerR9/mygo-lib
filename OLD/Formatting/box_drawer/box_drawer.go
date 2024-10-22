package box_drawer

import (
	"bytes"
	"io"
)

var (
	// DefaultBoxStyle is the default box style, that is, a padding of [1, 1, 1, 1] and
	// a line type of BtNormal (no heavy lines).
	DefaultBoxStyle *BoxStyle
)

func init() {
	DefaultBoxStyle = &BoxStyle{
		LineType: BtNormal,
		IsHeavy:  false,
		Padding:  [4]int{1, 1, 1, 1},
	}
}

// BoxBorderType is the type of the box border.
type BoxBorderType int

const (
	// BtNormal is the normal box border type.
	BtNormal BoxBorderType = iota

	// BtTriple is the triple box border type.
	BtTriple

	// BtQuadruple is the quadruple box border type.
	BtQuadruple

	// BtDouble is the double box border type.
	BtDouble

	// BtRounded is like BtNormal but with rounded corners.
	BtRounded
)

// BoxStyle is the style of the box.
type BoxStyle struct {
	// LineType is the type of the line.
	LineType BoxBorderType

	// IsHeavy is whether the line is heavy or not.
	// Only applicable to BtNormal, BtTriple, and BtQuadruple.
	IsHeavy bool

	// Padding is the padding of the box.
	// [Top, Right, Bottom, Left]
	Padding [4]int
}

// NewBoxStyle creates a new box style.
//
// Negative padding are set to 0.
//
// Parameters:
//   - line_type: The line type.
//   - is_heavy: Whether the line is heavy or not.
//   - padding: The padding of the box. [Top, Right, Bottom, Left]
//
// Returns:
//   - *BoxStyle: The new box style. Never returns nil.
func NewBoxStyle(line_type BoxBorderType, is_heavy bool, padding [4]int) *BoxStyle {
	for i := 0; i < 4; i++ {
		if padding[i] < 0 {
			padding[i] = 0
		}
	}

	bs := &BoxStyle{
		LineType: line_type,
		IsHeavy:  is_heavy,
		Padding:  padding,
	}

	return bs
}

var (
	// HeavyCorners is the heavy corners of the box.
	HeavyCorners [4][]byte

	// LightCorners is the light corners of the box.
	LightCorners [4][]byte
)

func init() {
	HeavyCorners = [4][]byte{
		[]byte("┏"), []byte("┓"), []byte("┗"), []byte("┛"),
	}

	LightCorners = [4][]byte{
		[]byte("┌"), []byte("┐"), []byte("└"), []byte("┘"),
	}
}

// Corners gets the corners of the box.
//
// Returns:
//   - [4][]byte: The corners. [TopLeft, TopRight, BottomLeft, BottomRight]
func (bs BoxStyle) Corners() [4][]byte {
	var corners [4][]byte

	if bs.IsHeavy {
		corners = HeavyCorners
	} else {
		corners = LightCorners
	}

	return corners
}

// TopBorder gets the top border of the box.
//
// It also applies to the bottom border as they are the same.
//
// Returns:
//   - string: The top border.
func (bs BoxStyle) TopBorder() []byte {
	var tb_border []byte

	switch bs.LineType {
	case BtNormal:
		if bs.IsHeavy {
			tb_border = []byte("━")
		} else {
			tb_border = []byte("─")
		}
	case BtTriple:
		if bs.IsHeavy {
			tb_border = []byte("┅")
		} else {
			tb_border = []byte("┄")
		}
	case BtQuadruple:
		if bs.IsHeavy {
			tb_border = []byte("┉")
		} else {
			tb_border = []byte("┅")
		}
	case BtDouble:
		tb_border = []byte("═")
	case BtRounded:
		tb_border = []byte("─")
	}

	return tb_border
}

// SideBorder gets the side border of the box.
//
// It also applies to the left border as they are the same.
//
// Returns:
//   - string: The side border.
func (bs BoxStyle) SideBorder() []byte {
	var side_border []byte

	switch bs.LineType {
	case BtNormal:
		if bs.IsHeavy {
			side_border = []byte("┃")
		} else {
			side_border = []byte("│")
		}
	case BtTriple:
		if bs.IsHeavy {
			side_border = []byte("┇")
		} else {
			side_border = []byte("┆")
		}
	case BtQuadruple:
		if bs.IsHeavy {
			side_border = []byte("┋")
		} else {
			side_border = []byte("┆")
		}
	case BtDouble:
		side_border = []byte("║")
	case BtRounded:
		side_border = []byte("│")
	}

	return side_border
}

// Apply draws a box around a content that is specified in a table.
//
// Format: If the content is [['H', 'e', 'l', 'l', 'o'], ['W', 'o', 'r', 'l', 'd']], the box will be:
//
//	┏━━━━━━━┓
//	┃ Hello ┃
//	┃ World ┃
//	┗━━━━━━━┛
//
// Parameters:
//   - table: The table that contains the content to be drawn.
//
// Returns:
//   - error: An error if the content could not be processed.
//
// Behaviors:
//   - If the box style is nil, the default box style will be used.
//
// Each string of the content represents a row in the box.
func (bs BoxStyle) Apply(w io.Writer, data []byte) (int, error) {
	for i := 0; i < 4; i++ {
		if bs.Padding[i] < 0 {
			bs.Padding[i] = 0
		}
	}

	left_padding := bytes.Repeat([]byte(" "), bs.Padding[3])
	right_padding := bytes.Repeat([]byte(" "), bs.Padding[1])

	tbb_char := bs.TopBorder()
	corners := bs.Corners()

	right_edge, err := rightMostEdge(data)
	if err != nil {
		return 0, err
	}

	lines := bytes.Split(data, []byte("\n"))

	lines = alignRightEdge(lines, right_edge)

	total_width := right_edge + bs.Padding[1] + bs.Padding[3]

	buff := NewWriter(w)

	err = buff.WriteMany(
		corners[0],
		bytes.Repeat(tbb_char, total_width),
		corners[1],
		Newline,
	)
	if err != nil {
		return buff.Written(), err
	}

	side_border := bs.SideBorder()

	for _, row := range lines {
		err := buff.WriteMany(
			side_border,
			left_padding,
			row,
			right_padding,
			side_border,
			Newline,
		)
		if err != nil {
			return buff.Written(), err
		}
	}

	err = buff.WriteMany(
		corners[2],
		bytes.Repeat(tbb_char, total_width),
		corners[3],
	)
	return buff.Written(), err
}
