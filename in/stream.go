package in

import (
	"errors"
	"unicode/utf8"

	"github.com/PlayerR9/mygo-lib/common"
)

// RuneScanner is a scanner of runes.
type RuneScanner struct {
	// data is the data to read from.
	data []byte

	// last_read is the last rune that was read.
	last_read []byte
}

// ReadRune implements io.RuneScanner.
func (rs *RuneScanner) ReadRune() (rune, int, error) {
	if rs == nil {
		return 0, 0, common.ErrNilReceiver
	}

	c, size := utf8.DecodeRune(rs.data)
	if c == utf8.RuneError {
		return 0, 0, errors.New("invalid UTF-8 sequence")
	}

	rs.last_read = rs.data[:size]
	rs.data = rs.data[size:]

	return c, size, nil
}

// Unread implements io.RuneScanner.
func (rs *RuneScanner) Unread() error {
	if rs == nil {
		return common.ErrNilReceiver
	}

	if len(rs.last_read) == 0 {
		return errors.New("no rune to unread")
	}

	rs.data = append(rs.last_read, rs.data...)
	rs.last_read = nil

	return nil
}

// Write implements io.Writer.
func (rs *RuneScanner) Write(data []byte) (int, error) {
	if rs == nil {
		return 0, common.ErrNilReceiver
	}

	rs.data = append(rs.data, data...)

	return len(data), nil
}
