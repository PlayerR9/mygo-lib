package strings

import (
	"github.com/PlayerR9/mygo-lib/strings/internal"
)

// ExtractFirstNFields extracts up to the first n fields from a given string.
//
// Parameters:
//   - str: The input string to split into fields.
//   - n: The maximum number of fields to extract.
//   - sep: The separator to use between fields.
//
// Returns:
//   - []string: A slice of fields extracted from the string, or nil if n is zero.
func ExtractFirstNFields(s *string, n uint, sep rune) []string {
	if s == nil || n == 0 {
		return nil
	}

	fn := func(r rune) bool {
		return r == sep
	}

	fields := internal.ExtractFirstNFields(s, n, fn)
	return fields
}

// ExtractFirstNFieldsFunc extracts up to the first n fields from a given string
// based on a separator function.
//
// Parameters:
//   - s: The input string pointer to split into fields.
//   - n: The maximum number of fields to extract.
//   - isSep: A function that determines if a rune is a separator.
//
// Returns:
//   - []string: A slice of fields extracted from the string, or nil if the input
//     string is nil, n is zero, or the separator function is nil.
//
// Panics:
//   - ErrNoPredicate: If the separator function is nil.
func ExtractFirstNFieldsFunc(s *string, n uint, isSep func(rune) bool) []string {
	if isSep == nil {
		panic(ErrNoPredicate)
	}

	if n == 0 || s == nil {
		return nil
	}

	fields := internal.ExtractFirstNFields(s, n, isSep)
	return fields
}
