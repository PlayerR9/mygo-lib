package common

// TODO panics with a TODO message. The given message is appended to the
// string "TODO: ". If the message is empty, the message "TODO: Handle this
// case" is used instead.
//
// Parameters:
//   - msg: The message to append to the string "TODO: ".
//
// This function is meant to be used only when the code is being built or
// refactored.
func TODO(msg string) {
	if msg == "" {
		panic("TODO: Handle this case")
	} else {
		panic("TODO: " + msg)
	}
}
