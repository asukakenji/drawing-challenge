package canvas

import "github.com/asukakenji/drawing-challenge/color"

// Canvas is a finite rectangular grid which allows several drawing operations.
// The coordinate system is zero-based.
type Canvas interface {
	// Dimensions returns the width and height.
	Dimensions() (int, int)

	// DrawLine draws a horizontal or vertical line.
	DrawLine(x1, y1, x2, y2 int) error

	// DrawRect draws a rectangle.
	DrawRect(x1, y1, x2, y2 int) error

	// BucketFill fills the area enclosing (x, y).
	BucketFill(x, y int, c color.Color) error
}

// BufferBasedCanvas is a Canvas based on a buffer of color.Color.
type BufferBasedCanvas interface {
	// Canvas is a super-interface of BufferBasedCanvas.
	Canvas

	// Set sets the color of the pixel at (x, y).
	Set(x, y int, c color.Color) error

	// At returns the color of the pixel at (x, y).
	At(x, y int) (color.Color, error)
}
