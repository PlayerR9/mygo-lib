package in

import (
	"io"
	"unicode/utf8"

	"github.com/PlayerR9/mygo-lib/common"
	flt "github.com/PlayerR9/mygo-lib/go-fault"
)

// runeRead is a rune reader.
type runeRead struct {
	// r is the reader.
	r io.Reader

	// stream is the rune stream.
	stream RuneStream

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
//   - flt.Fault: A fault if the provided reader is nil.
//
// Errors:
//   - common.ErrNilParam: If the input reader r is nil.
func privNewRuneRead(r io.Reader) (*runeRead, flt.Fault) {
	if r == nil {
		return nil, common.ErrNilParam("r")
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
		ok := r.stream.HasNext()
		if !ok {
			break
		}

		c, err := r.stream.Next()
		if err != nil {
			break
		}

		chars = append(chars, c)
	}

	r.chars = append(r.chars, chars...)

	return nil
}
