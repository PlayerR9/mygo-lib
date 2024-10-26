package generator

import (
	"bytes"
	"html/template"

	"github.com/PlayerR9/mygo-lib/common"
)

// Generater is an interface for data that holds information necessary for
// generating code.
type Generater interface {
	// SetPackageName sets the package name for the generator.
	//
	// Parameters:
	//   - pkg_name: The name of the package to be set.
	//
	// Returns:
	//   - error: An error if the receiver is nil.
	SetPackageName(pkg_name string) error
}

// Generate generates code from the given template and location.
//
// It first checks that the receiver, the template and the generator are not nil.
// Then it fixes the location to be a valid import path using the
// FixImportDir function. The fixed package name is then set to the generator
// using the SetPackageName method. After that, it executes the template using
// the generator as the data and returns the result as a byte slice.
//
// Parameters:
//   - templ: The template to generate code from.
//   - loc: The location of the package to be generated.
//   - gen: The generator to be used.
//
// Returns:
//   - []byte: The generated code.
//   - error: An error if any of the parameters are nil or if the package name
//     could not be fixed or if the template execution failed.
func Generate(templ *template.Template, loc string, gen Generater) ([]byte, error) {
	if templ == nil {
		return nil, common.NewErrNilParam("templ")
	} else if gen == nil {
		return nil, common.NewErrNilParam("gen")
	}

	pkg_name, err := FixImportDir(loc)
	if err != nil {
		return nil, err
	}

	err = gen.SetPackageName(pkg_name)
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer

	err = templ.Execute(&buff, gen)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
