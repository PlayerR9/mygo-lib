package common

// TODO panics with a TODO message. The given message is appended to the
// string "TODO: ". If the message is empty, the message "TODO: Handle this
// case" is used instead.
//
// Parameters:
//   - msg: The message to append to the string "TODO: ".
//
// This function is meant to be used only when the code is being built or
// refactored.
func TODO(msg string) {
	if msg == "" {
		panic("TODO: Handle this case")
	} else {
		panic("TODO: " + msg)
	}
}

// Pair is a pair of two values.
type Pair[A, B any] struct {
	// First is the first value.
	First A

	// Second is the second value.
	Second B
}

// NewPair creates a new pair.
//
// Parameters:
//   - first: The first value.
//   - second: The second value.
//
// Returns:
//   - Pair[A, B]: The new pair.
func NewPair[A, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{
		First:  first,
		Second: second,
	}
}
