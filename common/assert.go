package common

// FoundBug panics with the given message, indicating that a bug has been found.
//
// The message should be a single line of text explaining the bug.
//
// Parameters:
//   - msg: The message to be displayed when panicking.
//
// Panics with:
//   - [bug] <msg>
func FoundBug(msg string) {
	panic("[bug] " + msg)
}
