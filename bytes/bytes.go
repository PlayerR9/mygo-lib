package bytes

import (
	"io"

	"github.com/PlayerR9/go-errors/assert"
	"github.com/PlayerR9/mygo-lib/common"
)

type MultiWriter struct {
	w       io.Writer
	written int
}

func (w *MultiWriter) Write(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	} else if w == nil {
		return 0, common.ErrNilReceiver
	}

	n, err := w.w.Write(data)
	w.written += n

	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return n, err
}

func New(w io.Writer) (*MultiWriter, error) {
	if w == nil {
		return nil, common.NewErrNilParam("w")
	}

	return &MultiWriter{
		w: w,
	}, nil
}

func (w MultiWriter) Written() int {
	return w.written
}

func (w *MultiWriter) WriteBytes(data []byte) error {
	if len(data) == 0 {
		return nil
	} else if w == nil {
		return common.ErrNilReceiver
	}

	n, err := w.w.Write(data)
	w.written += n

	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

func (w *MultiWriter) WriteNewline() error {
	if w == nil {
		return common.ErrNilReceiver
	}

	n, err := w.w.Write(Newline)
	w.written += n

	if err == nil && n != NewlineLen {
		err = io.ErrShortWrite
	}

	return err
}

func (w *MultiWriter) WriteMany(datas ...[]byte) error {
	var total int

	for _, data := range datas {
		total += len(data)
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
