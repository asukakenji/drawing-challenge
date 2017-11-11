package bytecolor

import "github.com/asukakenji/drawing-challenge/color/bytecolor"

// point represents a point in the coordinate system.
// The coordinate system is zero-based.
type point struct {
	x int
	y int
}

// xyToIndex translates a 2D coordinate into a 1D index.
func xyToIndex(width, x, y int) int {
	return y*width + x
}

// isPointInsideCanvas returns whether (x, y) is within the bounds.
func isPointInsideCanvas(width, height, x, y int) bool {
	return 0 <= x && x < width && 0 <= y && y < height
}

// fill fills b with bc
// See the bytes.Repeat: https://golang.org/src/bytes/bytes.go
func fill(b []bytecolor.Color, bc bytecolor.Color) {
	b[0] = bc
	bp := 1
	for bp < len(b) {
		copy(b[bp:], b[:bp])
		bp *= 2
	}
}

// boolBuffer is a helper type for the bucket fill algorithm.
type boolBuffer struct {
	width  int
	height int
	values []bool
}

// newBoolBuffer returns a new boolBuffer.
func newBoolBuffer(width, height int) *boolBuffer {
	return &boolBuffer{
		width:  width,
		height: height,
		values: make([]bool, width*height),
	}
}

// At returns whether the pixel at (x, y) is processed.
// It returns true for any point outside the canvas
// to prevent it from really being processed.
func (bb *boolBuffer) At(x, y int) bool {
	index := xyToIndex(bb.width, x, y)
	if index < 0 || index >= len(bb.values) {
		return true
	}
	return bb.values[index]
}

// Set sets the pixel at (x, y) as already processed.
func (bb *boolBuffer) Set(x, y int) {
	index := xyToIndex(bb.width, x, y)
	bb.values[index] = true
}
