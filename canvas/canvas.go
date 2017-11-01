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
	Canvas

	// Set sets the color of the pixel at (x, y).
	Set(x, y int, c color.Color) error

	// At returns the color of the pixel at (x, y).
	At(x, y int) (color.Color, error)
}

// Point represetns a point in the coordinate system.
// The coordinate system is zero-based.
type Point struct {
	X int
	Y int
}

// NewPoint returns a new Point.
func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

// xyToIndex translates a 2D coordinate into a 1D index.
func xyToIndex(width, x, y int) int {
	return y*width + x
}

// isWithinCanvas returns whether (x, y) is within the bounds.
func isWithinCanvas(width, height, x, y int) bool {
	return 0 <= x && x < width && 0 <= y && y < height
}

// isHorizontalOrVertical returns
// whether (x1, y1) and (x2, y2) are horizontally or vertically aligned.
func isHorizontalOrVerticalLine(x1, y1, x2, y2 int) bool {
	return x1 == x2 || y1 == y2
}
