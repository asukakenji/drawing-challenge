package color

// Color represents a color value.
//
// There is no requirement on the underlying implementation
// as long as it is supported by the canvas which uses it.
type Color interface {
	// Equals returns whether this Color equals c.
	Equals(c Color) bool
}

// ByteColor represents a color value using a byte.
// It implements the Color interface.
type ByteColor byte

// Ensure that ByteColor implements the Color interface.
var (
	_ Color = ByteColor(0)
)

// Equals returns whether this Color equals c.
func (bc ByteColor) Equals(c Color) bool {
	bc2, ok := c.(ByteColor)
	if !ok {
		return false
	}
	return bc == bc2
}
