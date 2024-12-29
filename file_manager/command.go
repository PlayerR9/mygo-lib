package file_manager

import (
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// NewCommand creates a new command from the given name and arguments.
//
// Parameters:
//   - name: The name of the command to create.
//   - args: The arguments to the command.
//
// Returns:
//   - *exec.Cmd: The newly created command. Never returns nil.
//
// The command is created based on the current operating system.
//
// Panics:
//   - If the operating system is not supported. (i.e. Windows and Linux)
func NewCommand(name string, args ...string) *exec.Cmd {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		arg_line := strings.Join(args, " ")

		cmd = exec.Command(name, arg_line)
	case "linux":
		cmd = exec.Command(name, args...)
	default:
		panic("OS " + strconv.Quote(runtime.GOOS) + " is not supported")
	}

	return cmd
}
