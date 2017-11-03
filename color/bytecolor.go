package color

import "github.com/asukakenji/drawing-challenge/common"

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
