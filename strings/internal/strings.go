package internal

import "strings"

// ExtractFirstNFields extracts up to the first n fields from a given string.
//
// Parameters:
//   - str: The input string to split into fields.
//   - n: The maximum number of fields to extract.
//   - isSep: A function that returns true if the given rune is a separator.
//
// Returns:
//   - []string: A slice of fields extracted from the string, or nil if n is zero.
func ExtractFirstNFields(s *string, n uint, isSep func(rune) bool) []string {
	var result []string
	start := 0
	space_count := uint(0)

	for i, char := range *s {
		ok := isSep(char)

		if !ok {
			continue
		} else if space_count >= n {
			break
		}

		result = append(result, (*s)[start:i])
		space_count++
		start = i + 1
	}

	if space_count < n && start < len(*s) {
		result = append(result, (*s)[start:])
	}

	// Update the original string to remove the extracted fields
	if space_count < n {
		*s = ""
	} else {
		*s = (*s)[start:]
	}

	return result
}

// RejectEmpty rejects all empty strings in a slice of strings and returns
// the number of empty strings rejected.
//
// Parameters:
//   - s: The slice of strings to filter.
//
// Returns:
//   - uint: The number of empty strings rejected.
func RejectEmpty(s *[]string) uint {
	var end uint

	for _, str := range *s {
		if str == "" {
			continue
		}

		(*s)[end] = str
		end++
	}

	n := uint(len(*s)) - end

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return n
}

// Join concatenates a slice of strings into a single string with a separator.
//
// Parameters:
//   - s: The slice of strings to join.
//   - sep: The separator to use.
//
// Returns:
//   - string: The joined string.
//
// Example:
//
//	strs := []string{"a", "b", "c"}
//	strings.Join(&strs, ", ")
//	// returns "a, b, c"
func Join(s *[]string, sep string) string {
	_ = RejectEmpty(s)
	if len(*s) == 0 {
		return ""
	}

	var builder strings.Builder

	_, _ = builder.WriteString((*s)[0])

	for _, str := range (*s)[1:] {
		_, _ = builder.WriteString(sep)
		_, _ = builder.WriteString(str)
	}

	result := builder.String()
	return result
}
