package bytecolor

import (
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

// Parser parses a single-byte string to a Color.
// It implements the color.Parser interface.
type Parser struct {
	DefaultColor Color
}

// Ensure that Parser implements the color.Parser interface.
var (
	_ color.Parser = &Parser{}
)

// ParseColor parses s and returns a Color.
//
// Errors
//
// common.ErrInvalidColor:
// Will be returned if the color is not recognized by this parser.
//
func (parser *Parser) ParseColor(s string) (color.Color, error) {
	if s == "" {
		return parser.DefaultColor, nil
	}
	if len(s) != 1 {
		return Color(0), common.ErrInvalidColor
	}
	return Color(s[0]), nil
}
