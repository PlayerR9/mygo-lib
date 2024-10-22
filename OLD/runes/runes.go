package runes

import (
	"fmt"
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
)

// NormalizeRunes is a function that converts '\r\n' to '\n' and tabs to one or more spaces.
//
// Parameters:
//   - chars: The characters to convert.
//   - tab_size: The size of the tab.
//
// Returns:
//   - []rune: The normalized characters.
//   - error: An error if normalization failed.
//
// Errors:
//   - errors.ErrBadParameter: If tab_size is not positive.
//   - errors.ErrAfter: If the characters are not valid UTF-8.
func NormalizeRunes(chars []rune, tab_size int) ([]rune, error) {
	if tab_size <= 0 {
		return chars, common.NewErrBadParameter("tab_size", fmt.Sprintf("must be positive, got %d", tab_size))
	}

	if len(chars) == 0 {
		return nil, nil
	}

	indices := IndicesOf(chars, '\r')

	for _, idx := range indices {
		if idx+1 >= len(chars) {
			return chars, NewErrAfter(chars[idx], NewErrNotAsExpected("rune", nil, '\n'))
		}

		next := chars[idx+1]
		if next != '\n' {
			return chars, NewErrAfter(chars[idx], NewErrNotAsExpected("rune", &next, '\n'))
		}
	}

	for i, idx := range indices {
		chars = slices.Delete(chars, idx-i, idx-i+1)
	}

	var count int

	for i := 0; i < len(chars); i++ {
		if chars[i] == '\t' {
			count++
		}
	}

	if count == 0 {
		return chars, nil
	}

	new_chars := make([]rune, 0, len(chars)+count*(tab_size-1))

	tab := make([]rune, 0, tab_size)
	for i := 0; i < tab_size; i++ {
		tab = append(tab, ' ')
	}

	for _, c := range chars {
		if c == '\t' {
			new_chars = append(new_chars, tab...)
		} else {
			new_chars = append(new_chars, c)
		}
	}

	return new_chars, nil
}
