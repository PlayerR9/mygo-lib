package generator

import (
	"errors"
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
)

// FixImportDir fixes the given location to be a valid import path.
//
// It first checks that the location is not empty. Then it checks that the
// extension is ".go". After that, it checks that the directory of the location
// is not the current directory. If it is, it gets the package name using the
// "go build" package and returns it. If not, it splits the directory of the
// location and returns the second element of the split.
//
// If any of the above checks fails, it returns an error.
//
// Parameters:
//   - loc: The location to be fixed.
//
// Returns:
//   - string: The fixed import path.
//   - error: An error if the location could not be fixed.
func FixImportDir(loc string) (string, error) {
	if loc == "" {
		return "", errors.New("empty location")
	}

	ext := filepath.Ext(loc)
	if ext != ".go" {
		return "", fmt.Errorf("expected %q extension, got %q", ".go", ext)
	}

	dir_loc := filepath.Dir(loc)
	if dir_loc != "." {
		_, dir := filepath.Split(dir_loc)
		return dir, nil
	}

	pkg, err := build.ImportDir(loc, 0)
	if err != nil {
		return "", err
	}

	return pkg.Name, nil
}

// NewLogger creates a new logger with the given name.
//
// If the name is empty, it defaults to "generator".
//
// Parameters:
//   - name: The name of the logger.
//
// Returns:
//   - *log.Logger: The logger. Never returns nil.
func NewLogger(name string) *log.Logger {
	if name == "" {
		name = "generator"
	}

	return log.New(os.Stdout, "["+name+"]: ", log.Lshortfile)
}
