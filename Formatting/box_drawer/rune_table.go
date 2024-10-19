package box_drawer

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func RightMostEdge(table []byte) (int, error) {
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

func AlignRightEdge(table [][]byte, edge int) [][]byte {
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
