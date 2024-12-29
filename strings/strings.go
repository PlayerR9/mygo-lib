package strings

import (
	"github.com/PlayerR9/mygo-lib/strings/internal"
)

// ExtractFirstNFields extracts up to the first n fields from a given string.
//
// Parameters:
//   - str: The input string to split into fields.
//   - n: The maximum number of fields to extract.
//
// Returns:
//   - []string: A slice of fields extracted from the string, or nil if n is zero.
func ExtractFirstNFields(s *string, n uint) []string {
	if s == nil || n == 0 {
		return nil
	}

	fields := internal.ExtractFirstNFields(s, n)
	return fields
}
