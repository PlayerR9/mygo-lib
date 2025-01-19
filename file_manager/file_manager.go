package file_manager

import (
	"errors"
	"os"
)

// Exists checks if the given location exists.
//
// Parameters:
//   - loc: The location to check.
//
// Returns:
//   - bool: True if the location exists, false otherwise.
//   - error: An error if something went wrong.
func Exists(loc string) (bool, error) {
	_, err := os.Stat(loc)
	if err == nil {
		return true, nil
	}

	ok := errors.Is(err, os.ErrNotExist)
	if !ok {
		return false, err
	}

	return false, nil
}

// CreateDirectory creates a directory at the given location with the given mode.
//
// Parameters:
//   - loc: The location to create the directory.
//   - mode: The file mode to use when creating the directory.
//   - force: If true, the directory will be overwritten if it already exists.
//
// Returns:
//   - error: An error if the directory cannot be created.
//
// Errors:
//   - os.ErrExist: If the directory already exists and force is false.
//   - any other error: If the directory cannot be created.
func CreateDirectory(loc string, mode os.FileMode, force bool) error {
	_, err := os.Stat(loc)
	if err != nil {
		ok := errors.Is(err, os.ErrNotExist)
		if !ok {
			return err
		}

		if err := os.Mkdir(loc, mode); err != nil {
			return err
		}

		return nil
	}

	if !force {
		return os.ErrExist
	}

	if err := os.RemoveAll(loc); err != nil {
		return err
	}

	if err := os.Mkdir(loc, mode); err != nil {
		return err
	}

	return nil
}
