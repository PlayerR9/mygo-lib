package in

import (
	"errors"
	"io"
	"unicode/utf8"
)

type runeRead struct {
	r      io.Reader
	stream RuneStream
	chars  []rune
}

func newRuneRead(r io.Reader) (*runeRead, error) {
	if r == nil {
		return nil, errors.New("parameter (r) must not be nil")
	}

	return &runeRead{
		r: r,
	}, nil
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
