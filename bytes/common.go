package bytes

// Newline is a newline character.
var Newline []byte

// NewlineLen is the length of Newline.
const NewlineLen uint = 1

func init() {
	Newline = []byte{'\n'}
}
