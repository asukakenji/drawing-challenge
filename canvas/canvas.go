package canvas

import "github.com/asukakenji/drawing-challenge/color"

// Canvas is a finite rectangular grid which allows several drawing operations.
// The coordinate system is zero-based.
type Canvas interface {
	// Dimensions returns the width and height.
	Dimensions() (int, int)

	// DrawLine draws a horizontal or vertical line.
	//
	// Errors
	//
	// common.ErrPointOutsideCanvas:
	// Will be returned if (x1, y1) or (x2, y2) is outside the canvas.
	//
	// common.ErrLineNotHorizontalOrVertical:
	// Will be returned if the line is not horizontal or vertical.
	//
	DrawLine(x1, y1, x2, y2 int) error

	// DrawRect draws a rectangle.
	//
	// Errors
	//
	// common.ErrPointOutsideCanvas:
	// Will be returned if (x1, y1) or (x2, y2) is outside the canvas.
	//
	DrawRect(x1, y1, x2, y2 int) error

	// BucketFill fills the area enclosing (x, y). The pixels connecting to
	// (x, y) having the same color as that at (x, y) are replaced by c.
	//
	// Errors
	//
	// common.ErrPointOutsideCanvas:
	// Will be returned if (x, y) is outside the canvas.
	//
	// common.ErrColorTypeNotSupported:
	// Will be returned if c is not supported by the canvas.
	//
	BucketFill(x, y int, c color.Color) error
}

// BufferBasedCanvas is a Canvas based on a buffer of color.Color.
type BufferBasedCanvas interface {
	// Canvas is a super-interface of BufferBasedCanvas.
	Canvas

	// At returns the color of the pixel at (x, y).
	//
	// Errors
	//
	// common.ErrPointOutsideCanvas:
	// Will be returned if (x, y) is outside the canvas.
	//
	At(x, y int) (color.Color, error)

	// Set sets the color of the pixel at (x, y).
	//
	// Errors
	//
	// common.ErrPointOutsideCanvas:
	// Will be returned if (x, y) is outside the canvas.
	//
	// common.ErrColorTypeNotSupported:
	// Will be returned if c is not supported by the canvas.
	//
	Set(x, y int, c color.Color) error
}
