package display

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	gctx "github.com/PlayerR9/mygo-lib/Contexts"
	fbd "github.com/PlayerR9/mygo-lib/Formatting/box_drawer"
)

// debugAct is an action that prints debug information to the logger.
type debugAct struct {
	// text is the text to print.
	text string
}

// Run implements the contexts.Action interface.
func (act debugAct) Run(ctx context.Context) error {
	c, err := fromContext(ctx)
	if err != nil {
		return err
	}

	c.logger.Println(act.text)

	return nil
}

// Debug returns an action that prints the given text to the logger as a debug
// message. If the text is empty, nil is returned.
//
// Parameters:
//   - text: The text to print.
//
// Returns:
//   - contexts.Action: An action that prints the given text to the logger as a
func Debug(text string) gctx.Action {
	if text == "" {
		return nil
	}

	return &debugAct{
		text: text,
	}
}

// Debugf returns an action that prints the given formatted string to the logger
// as a debug message. If the formatted string is empty, nil is returned.
//
// Parameters:
//   - format: The format string.
//   - args: The arguments to format.
//
// Returns:
//   - contexts.Action: An action that prints the given formatted string to the
//     logger as a debug message. If the formatted string is empty, nil is
//     returned.
func Debugf(format string, args ...any) gctx.Action {
	str := fmt.Sprintf(format, args...)
	if str == "" {
		return nil
	}

	return &debugAct{
		text: str,
	}
}

// normalAct is an action that prints data to the logger.
type normalAct struct {
	// data is the data to print.
	data []byte
}

// Run implements the contexts.Action interface.
func (act normalAct) Run(ctx context.Context) error {
	c, err := fromContext(ctx)
	if err != nil {
		return err
	}

	w := c.logger.Writer()

	n, err := w.Write(act.data)
	if err == nil && n != len(act.data) {
		err = io.ErrShortWrite
	}

	return err
}

// Write returns an action that writes the given data to the logger followed by a
// newline character. If the data is empty, nil is returned.
//
// Parameters:
//   - data: The data to write.
//
// Returns:
//   - contexts.Action: An action that writes the given data to the logger
//     followed by a newline character. If the data is empty, nil is returned.
func Write(data []byte) gctx.Action {
	data = append(data, '\n')

	return &normalAct{
		data: data,
	}
}

// Writef returns an action that writes the formatted string to the logger
// followed by a newline character.
//
// Parameters:
//   - format: The format string.
//   - args: The arguments to format.
//
// Returns:
//   - contexts.Action: An action that writes the formatted string to the
//     logger followed by a newline character.
func Writef(format string, args ...any) gctx.Action {
	return &normalAct{
		data: append([]byte(fmt.Sprintf(format, args...)), '\n'),
	}
}

// Newline returns an action that writes a newline character to the logger.
//
// Returns:
//   - contexts.Action: An action that writes a newline character to the logger.
func Newline() gctx.Action {
	return &normalAct{
		data: []byte{'\n'},
	}
}

// boxedAct is an action that writes data to the logger in a box.
type boxedAct struct {
	// data is the data to write.
	data []byte
}

// Run implements the contexts.Action interface.
func (act boxedAct) Run(ctx context.Context) error {
	c, err := fromContext(ctx)
	if err != nil {
		return err
	}

	box := fbd.NewBoxStyle(fbd.BtNormal, true, [4]int{0, 1, 0, 1})

	var buff bytes.Buffer

	total, err := box.Apply(&buff, act.data)
	if err != nil {
		return err
	}

	w := c.logger.Writer()

	n, err := w.Write(buff.Bytes())
	if err == nil && n != total {
		err = io.ErrShortWrite
	}

	return err
}

// Boxed returns an action that writes the given data to the logger in a box.
//
// Parameters:
//   - data: The data to write.
//
// Returns:
//   - contexts.Action: An action that writes the given data to the logger in a box.
func Boxed(data []byte) gctx.Action {
	if data == nil {
		return nil
	}

	return &boxedAct{
		data: data,
	}
}

// waitEnterAct waits for the user to press enter.
type waitEnterAct struct{}

// Run implements the contexts.Action interface.
func (act waitEnterAct) Run(ctx context.Context) error {
	c, err := fromContext(ctx)
	if err != nil {
		return err
	}

	w := c.logger.Writer()

	data := []byte("Press ENTER to continue...\n")

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	fmt.Scanln()

	return nil
}

// WaitEnter returns an action that waits for the user to press enter.
//
// Returns:
//   - gctx.Action: An action that prompts the user to press ENTER to continue.
func WaitEnter() gctx.Action {
	return &waitEnterAct{}
}

// clearAct clears the buffer in the context.
type clearAct struct{}

// Run implements the contexts.Action interface.
func (act clearAct) Run(ctx context.Context) error {
	c, err := fromContext(ctx)
	if err != nil {
		return err
	}

	w := c.logger.Writer()

	if _, err := w.Write([]byte("\x1B[2J\x1B[H")); err != nil {
		return err
	}

	return nil
}

// Clear returns an action that clears the buffer in the context.
//
// Returns:
//   - gctx.Action: An action that clears the buffer in the context. Never returns nil.
func Clear() gctx.Action {
	return &clearAct{}
}

// waitAct waits for the given number of milliseconds.
type waitAct struct {
	ms int
}

// Run implements the contexts.Action interface.
func (act waitAct) Run(ctx context.Context) error {
	time.Sleep(time.Duration(act.ms) * time.Millisecond)

	return nil
}

// Pause returns an action that waits for the given number of milliseconds.
//
// Parameters:
//   - ms: The number of milliseconds to wait.
//
// Returns:
//   - gctx.Action: An action that waits for the given number of milliseconds.
//     If ms is 0 or less, returns nil.
func Pause(ms int) gctx.Action {
	if ms <= 0 {
		return nil
	}

	return &waitAct{
		ms: ms,
	}
}
