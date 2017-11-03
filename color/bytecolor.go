package color

import "github.com/asukakenji/drawing-challenge/common"

// ByteColor represents a color value using a byte.
// It implements the Color interface.
type ByteColor byte

// Ensure that ByteColor implements the Color interface.
var (
	_ Color = ByteColor(0)
)

// ParseByteColor returns a ByteColor from a single-byte string.
//
// Errors
//
// common.ErrEmptyColor:
// Will be returned if s is an empty string.
//
// common.ErrInvalidColor:
// Will be returned if s is not a single-byte string.
//
func ParseByteColor(s string) (ByteColor, error) {
	if s == "" {
		return ByteColor(0), common.ErrEmptyColor
	}
	if len(s) != 1 {
		return ByteColor(0), common.ErrInvalidColor
	}
	return ByteColor(s[0]), nil
}

// Equals returns whether this Color equals c.
func (bc ByteColor) Equals(c Color) bool {
	bc2, ok := c.(ByteColor)
	if !ok {
		return false
	}
	return bc == bc2
}
