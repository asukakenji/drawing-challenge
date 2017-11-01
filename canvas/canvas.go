package canvas

import "github.com/asukakenji/drawing-challenge/color"

// Canvas TODO
type Canvas interface {
	Dimensions() (int, int)
	DrawLine(x1, y1, x2, y2 int) error
	DrawRect(x1, y1, x2, y2 int) error
	BucketFill(x, y int, c color.Color) error
}

// BufferBasedCanvas TODO
type BufferBasedCanvas interface {
	Canvas
	Set(x, y int, c color.Color) error
	At(x, y int) (color.Color, error)
}

// Point TODO
type Point struct {
	X int
	Y int
}

// NewPoint TODO
func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

// xyToIndex TODO
func xyToIndex(width, x, y int) int {
	return y*width + x
}

// isWithinCanvas TODO
func isWithinCanvas(w, h, x, y int) bool {
	return 0 <= x && x < w && 0 <= y && y < h
}

// isHorizontalOrVertical TODO
func isHorizontalOrVerticalLine(x1, y1, x2, y2 int) bool {
	return x1 == x2 || y1 == y2
}
