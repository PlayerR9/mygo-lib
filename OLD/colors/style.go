package colors

import "bytes"

var (
	// DefaultStyle is the default style. It is white foreground and black background.
	DefaultStyle *Style
)

func init() {
	DefaultStyle = &Style{
		fg: White,
		bg: Black,
	}
}

// Style represents a style.
type Style struct {
	// fg is the foreground color.
	fg Color

	// bg is the background color.
	bg Color
}

// String implements the fmt.Stringer interface.
func (s Style) String() string {
	var buff bytes.Buffer

	_ = s.fg.Foreground(&buff)
	_ = s.bg.Background(&buff)

	return buff.String()
}

// Foreground sets the foreground color of the style.
//
// Parameters:
//   - color: The color to set as the foreground.
//
// Returns:
//   - *Style: A new Style with the updated foreground color. Never returns nil.
func (s *Style) Foreground(color Color) *Style {
	if s == nil {
		return &Style{
			fg: color,
			bg: Black,
		}
	} else {
		return &Style{
			fg: color,
			bg: s.bg,
		}
	}
}

// Background sets the background color of the style.
//
// Parameters:
//   - color: The color to set as the background.
//
// Returns:
//   - *Style: A new Style with the updated background color. Never returns nil.
func (s *Style) Background(color Color) *Style {
	if s == nil {
		return &Style{
			fg: White,
			bg: color,
		}
	} else {
		return &Style{
			fg: s.fg,
			bg: color,
		}
	}
}