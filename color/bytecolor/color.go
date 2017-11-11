// Package bytecolor defines the ByteColor type,
// which implements the color.Color interface,
// and the Parser type,
// which implements the color.Parser interface.
package bytecolor

import "github.com/asukakenji/drawing-challenge/color"

// Color represents a color value using a byte.
// It implements the color.Color interface.
type Color byte

// Ensure that Color implements the color.Color interface.
var (
	_ color.Color = Color(0)
)

// Equals returns whether this Color equals c.
func (bc Color) Equals(c color.Color) bool {
	bc2, ok := c.(Color)
	if !ok {
		return false
	}
	return bc == bc2
}
