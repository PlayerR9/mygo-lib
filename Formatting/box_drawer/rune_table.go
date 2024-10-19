package box_drawer

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// rightMostEdge calculates the length of the longest line in the given byte slice table.
// It iterates through the UTF-8 encoded bytes to determine the longest line length.
// If an invalid UTF-8 character is encountered, it returns an error with details.
//
// Parameters:
//   - table: The byte slice containing the UTF-8 encoded text.
//
// Returns:
//   - int: The length of the longest line.
//   - error: An error if an invalid UTF-8 character is encountered.
func rightMostEdge(table []byte) (int, error) {
	if len(table) == 0 {
		return 0, nil
	}

	var longest_line, current int

	for len(table) > 0 {
		r, size := utf8.DecodeRune(table)
		table = table[size:]

		switch r {
		case utf8.RuneError:
			return longest_line, fmt.Errorf("invalid UTF-8: %q", table)
		case '\n':
			if current > longest_line {
				longest_line = current
			}

			current = 0
		default:
			current++
		}
	}

	if current > longest_line {
		longest_line = current
	}

	return longest_line, nil
}

// Aligns the right edge of a table to a certain width.
//
// This function appends spaces to the end of each row in the table until the
// row is at least `edge` characters wide. It does not modify the contents of
// the table, but instead returns a new table with the modified rows.
//
// The returned table is a copy of the original table, so the original table
// is not modified.
//
// Parameters:
//   - table: The table to align.
//   - edge: The width to align to.
//
// Returns:
//   - [][]byte: The aligned table.
func alignRightEdge(table [][]byte, edge int) [][]byte {
	if len(table) == 0 {
		return table
	}

	for i := 0; i < len(table); i++ {
		curr_row := table[i]
		curr_size := utf8.RuneCount(curr_row)

		padding := edge - curr_size

		padding_right := bytes.Repeat([]byte{' '}, padding)

		table[i] = append(curr_row, padding_right...)
	}

	return table
}
