package color

import "github.com/asukakenji/drawing-challenge/common"

// Parser represents a color parser.
type Parser interface {
	// ParseColor parses s and returns a Color.
	//
	// Errors
	//
	// common.ErrInvalidColor:
	// Will be returned if the color is not recognized by this parser.
	//
	ParseColor(s string) (Color, error)
}

// ByteColorParser parses a single-byte string to a ByteColor.
// It implements the Parser interface.
type ByteColorParser struct {
	DefaultColor ByteColor
}

// Ensure that ByteColor implements the Color interface.
var (
	_ Parser = &ByteColorParser{}
)

// ParseColor parses s and returns a Color.
//
// Errors
//
// common.ErrInvalidColor:
// Will be returned if the color is not recognized by this parser.
//
func (parser *ByteColorParser) ParseColor(s string) (Color, error) {
	if s == "" {
		return parser.DefaultColor, nil
	}
	if len(s) != 1 {
		return ByteColor(0), common.ErrInvalidColor
	}
	return ByteColor(s[0]), nil
}
