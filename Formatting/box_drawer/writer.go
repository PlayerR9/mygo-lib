package box_drawer

import (
	"io"

	assert "github.com/PlayerR9/go-verify"
	gers "github.com/PlayerR9/mygo-lib/errors"
)

var (
	// Newline is a newline character.
	Newline []byte
)

const NewlineLen int = 1

func init() {
	Newline = []byte{'\n'}
}

// Writer is a wrapper around an io.Writer that keeps track of how much was written.
type Writer struct {
	// w is the underlying io.Writer.
	w io.Writer

	// written is the number of bytes written.
	written int
}

// NewWriter creates a new Writer. If w is nil, io.Discard is used.
//
// Parameters:
//   - w: The underlying io.Writer.
//
// Returns:
//   - Writer: The new Writer.
func NewWriter(w io.Writer) Writer {
	if w == nil {
		w = io.Discard
	}

	return Writer{
		w: w,
	}
}

// Written returns the number of bytes written.
//
// Returns:
//   - int: The number of bytes written.
func (w Writer) Written() int {
	return w.written
}

// WriteBytes writes the data to the underlying io.Writer.
//
// Parameters:
//   - data: The data to write.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *Writer) WriteBytes(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if w == nil {
		return gers.ErrNilReceiver
	}

	assert.Cond(w.w != nil, "w must not be nil")

	n, err := w.w.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	w.written += n

	return err
}

// WriteNewline writes a newline character to the underlying io.Writer.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *Writer) WriteNewline() error {
	if w == nil {
		return gers.ErrNilReceiver
	}

	assert.Cond(w.w != nil, "w must not be nil")

	n, err := w.w.Write(Newline)
	if err == nil && n != NewlineLen {
		err = io.ErrShortWrite
	}

	w.written += n

	return err
}

// WriteMany writes many data to the underlying io.Writer.
//
// Parameters:
//   - datas: The datas to write.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *Writer) WriteMany(datas ...[]byte) error {
	var total int

	for _, data := range datas {
		total += len(data)
	}

	if total == 0 {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	assert.Cond(w.w != nil, "w must not be nil")

	final := make([]byte, total)
	var prev int

	for _, data := range datas {
		copy(final[prev:], data)
		prev += len(data)
	}

	err := w.WriteBytes(final)
	return err
}
