package color

// Color represents a color value.
// There is no requirement on the underlying implementation as far as
// it is understood by the Canvas which uses it.
type Color interface {
	// Equals returns whether this Color equals c.
	Equals(c Color) bool
}
