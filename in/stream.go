package in

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type RuneStream struct {
	data []byte
}

func (rs *RuneStream) Write(data []byte) (int, error) {
	if rs == nil {
		return 0, errors.New("receiver must not be nil")
	}

	rs.data = append(rs.data, data...)

	return len(data), nil
}

func (rs *RuneStream) Next() (rune, error) {
	if rs == nil {
		return 0, errors.New("receiver must not be nil")
	}

	c, size := utf8.DecodeRune(rs.data)
	if c == utf8.RuneError {
		return 0, fmt.Errorf("invalid UTF-8 sequence")
	}

	rs.data = rs.data[size:]

	return c, nil
}

func (rs RuneStream) HasNext() bool {
	return len(rs.data) > 0
}
