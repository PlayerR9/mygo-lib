package internal

import (
	"io"

	"github.com/PlayerR9/mygo-lib/common"
)

// MakeWTInfo is used by WriteTree to print the token tree in a form that is easier to read
// and understand.
type MakeWTInfo struct {
	// w is the writer to write to.
	w io.Writer

	// n is the number of bytes written.
	n *int

	// indent is the current indentation.
	indent []byte

	// has_next is true if there is a next sibling, false otherwise.
	has_next bool

	// is_first is true if this is the root node, false otherwise.
	is_first bool
}

/* // Validate returns an error if the receiver's inner state is invalid.
//
// Returns:
//   - error: An error if the receiver is invalid.
//
// The receiver is valid if the following conditions are met:
//   - The writer is not nil.
//   - The number of bytes written is not nil.
//   - The number of bytes written is not negative.
func (info MakeWTInfo) Validate() error {
	if info.w == nil {
		return errors.New("info.w must not be nil")
	}

	if info.n == nil {
		return errors.New("info.n must not be nil")
	} else if *info.n < 0 {
		return errors.New("info.n must not be negative")
	}

	return nil
} */

// NewMakeWTInfo returns a new instance of MakeWTInfo.
//
// Parameters:
//   - w: The writer to write to.
//
// Returns:
//   - *MakeWTInfo: A new instance of MakeWTInfo. Never returns nil.
func NewMakeWTInfo(w io.Writer) *MakeWTInfo {
	if w == nil {
		w = io.Discard
	}

	var n int

	return &MakeWTInfo{
		w:        w,
		n:        &n,
		indent:   nil,
		has_next: false,
		is_first: true,
	}
}

// New returns a new instance of MakeWTInfo with the given indent.
//
// Parameters:
//   - indent: The indent to use.
//
// Returns:
//   - *MakeWTInfo: A new instance of MakeWTInfo. Never returns nil.
func (info MakeWTInfo) New(indent []byte) *MakeWTInfo {
	return &MakeWTInfo{
		w:        info.w,
		n:        info.n,
		indent:   indent,
		has_next: true,
		is_first: false,
	}
}

// WriteData writes the given data to the writer.
//
// Parameters:
//   - data: The data to write.
//   - expected_size: The expected size of the data.
//
// Returns:
//   - error: An error if the data could not be written. Returns nil if the
//     data was written successfully.
//
// Errors:
//   - io.ErrShortWrite: If the data was not fully written.
//   - errors.ErrNilReceiver: If the receiver is nil.
//   - any other error: If an error occurs while writing the data.
func (info *MakeWTInfo) WriteData(data []byte, expected_size int) error {
	if info == nil {
		return common.ErrNilReceiver
	}

	// err := info.Validate()
	// assert.Err(err, "info.Validate()")

	if len(data) == 0 {
		if expected_size != len(data) {
			return io.ErrShortWrite
		}

		return nil
	}

	n, err := info.w.Write(data)
	if err == nil && n != expected_size {
		err = io.ErrShortWrite
	}

	*info.n += n

	return err
}

// WriteString is like WriteData but writes the given string to the writer.
//
// Parameters:
//   - str: The string to write.
//
// Returns:
//   - error: An error if the string could not be written. Returns nil if the
//     string was written successfully.
//
// Errors:
//   - io.ErrShortWrite: If the string was not fully written.
//   - errors.ErrNilReceiver: If the receiver is nil.
//   - any other error: If an error occurs while writing the string.
func (info *MakeWTInfo) WriteString(str string) error {
	if info == nil {
		return common.ErrNilReceiver
	}

	// err := info.Validate()
	// assert.Err(err, "info.Validate()")

	if str == "" {
		return nil
	} else if info.w == nil {
		return io.ErrShortWrite
	}

	data := []byte(str)

	n, err := info.w.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	*info.n += n

	return err
}

// IsFirst checks whether this is the first node in the tree.
//
// Returns:
//   - bool: True if this is the first node in the tree, false otherwise.
func (info MakeWTInfo) IsFirst() bool {
	return info.is_first
}

// IsLast checks whether this is the last node in the tree.
//
// Returns:
//   - bool: True if this is the last node in the tree, false otherwise.
func (info MakeWTInfo) IsLast() bool {
	return !info.has_next
}

// SetIsLast sets whether this is the last node in the tree.
//
// Parameters:
//   - is_last: True if this is the last node in the tree, false otherwise.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (info *MakeWTInfo) SetIsLast(is_last bool) bool {
	if info == nil {
		return false
	}

	info.has_next = !is_last

	return true
}

// Indent returns the current indent.
//
// Returns:
//   - []byte: The current indent.
func (info MakeWTInfo) Indent() []byte {
	return info.indent
}

// BytesWritten returns the number of bytes written.
//
// Returns:
//   - int: The number of bytes written.
func (info MakeWTInfo) BytesWritten() int {
	// err := info.Validate()
	// assert.Err(err, "info.Validate()")

	return *info.n
}
