package in

import (
	"io"
	"unicode/utf8"

	"github.com/PlayerR9/mygo-lib/common"
)

// runeRead is a rune reader.
type runeRead struct {
	// r is the reader.
	r io.Reader

	// stream is the rune stream.
	stream RuneScanner

	// chars is the current chars.
	chars []rune
}

// privNewRuneRead initializes a new runeRead instance with the provided io.Reader.
//
// Parameters:
//   - r: The io.Reader to read from.
//
// Returns:
//   - *runeRead: A pointer to the newly created runeRead instance.
//   - error: A fault if the provided reader is nil.
//
// Errors:
//   - common.ErrBadParam: If the input reader r is nil.
func privNewRuneRead(r io.Reader) (*runeRead, error) {
	if r == nil {
		return nil, common.NewErrNilParam("r")
	}

	rr := &runeRead{
		r: r,
	}

	return rr, nil
}

func (r *runeRead) read() error {
	var data [utf8.UTFMax]byte

	n, err := r.r.Read(data[:])
	if err != nil {
		return err
	}

	_, _ = r.stream.Write(data[:n])

	return nil
}

func (r *runeRead) getRunes() error {
	err := r.read()
	if err != nil {
		return err
	}

	var chars []rune

	for {
		c, _, err := r.stream.ReadRune()
		if err != nil {
			break
		}

		chars = append(chars, c)
	}

	r.chars = append(r.chars, chars...)

	return nil
}
