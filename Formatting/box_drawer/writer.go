package box_drawer

import (
	"io"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

var (
	Newline    []byte
	NewlineLen int
)

func init() {
	Newline = []byte("\n")
	NewlineLen = len(Newline)
}

type Writer struct {
	w       io.Writer
	written int
}

func NewWriter(w io.Writer) Writer {
	if w == nil {
		w = io.Discard
	}

	return Writer{
		w: w,
	}
}

func (w Writer) Written() int {
	return w.written
}

func (w *Writer) WriteBytes(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if w == nil {
		return gers.ErrNilReceiver
	}

	n, err := w.w.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	w.written += n

	return err
}

func (w *Writer) WriteNewline() error {
	if w == nil {
		return gers.ErrNilReceiver
	}

	n, err := w.w.Write(Newline)
	if err == nil && n != NewlineLen {
		err = io.ErrShortWrite
	}

	w.written += n

	return err
}

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

	final := make([]byte, total)
	var prev int

	for _, data := range datas {
		copy(final[prev:], data)
		prev += len(data)
	}

	err := w.WriteBytes(final)
	return err
}
