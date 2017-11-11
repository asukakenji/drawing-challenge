// Package color defines the Color interface and the Parser interface.
package color

// Color represents a color value.
//
// There is no requirement on the underlying implementation
// as long as it is supported by the canvas which uses it.
type Color interface {
	// Equals returns whether this Color equals c.
	Equals(c Color) bool
}

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
