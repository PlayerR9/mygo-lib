package in

import (
	"io"
	"os"

	"github.com/PlayerR9/mygo-lib/optional"
	gslc "github.com/PlayerR9/mygo-lib/slices"
)

func ScanLn() ([]rune, error) {
	rr, err := privNewRuneRead(os.Stdin)
	if err != nil {
		panic(err)
	}

	var input []rune

	for {
		err := rr.getRunes()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		result := gslc.FirstIndexOf(rr.chars, '\n')
		if !result.IsPresent() {
			input = append(input, rr.chars...)
			rr.chars = nil

			continue
		}

		idx := optional.MustGet[uint](result)
		if idx != 0 {
			input = append(input, rr.chars[:idx-1]...)
		} else if len(input) > 0 && input[len(input)-1] == '\r' {
			input = input[:len(input)-1]
		}

		rr.chars = rr.chars[idx+1:]

		break
	}

	return input, nil
}
