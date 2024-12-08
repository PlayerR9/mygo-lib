package bytes

import (
	"unicode/utf8"

	"github.com/PlayerR9/mygo-lib/common"
)

// Encode appends the UTF-8 encoding of a slice of runes to the provided byte slice.
//
// Parameters:
//   - data: A pointer to a byte slice where the encoded UTF-8 bytes will be appended.
//   - chars: A slice of runes to encode.
//
// Panics:
//   - common.ErrBadParam: If data is nil.
func Encode(data *[]byte, chars []rune) {
	if data == nil {
		panic(common.NewErrNilParam("data"))
	}

	for _, c := range chars {
		*data = utf8.AppendRune(*data, c)
	}
}
