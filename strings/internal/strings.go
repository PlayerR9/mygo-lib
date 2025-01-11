package internal

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
