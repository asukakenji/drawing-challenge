package color

// Color represents a color value.
//
// There is no requirement on the underlying implementation as long as
// it is supported by the canvas which uses it.
type Color interface {
	// Equals returns whether this Color equals c.
	Equals(c Color) bool
}
