package strings

import "strings"

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

	// TODO: Modify here such that this function does not depend on
	// "strings.Split" method of the "strings" package.
	fields := strings.Split(*s, " ")

	var min int

	if len(fields) < int(n) {
		min = len(fields)
	} else {
		min = int(n)
	}

	*s = strings.Join(fields[min:], " ")

	result := make([]string, 0, n)

	for i := 0; i < min; i++ {
		result = append(result, fields[i])
	}

	result = result[:len(result):len(result)]

	return result
}
