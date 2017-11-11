// Package bytecolor defines the Buffer type,
// which implements the canvas.BufferBasedCanvas interface.
package bytecolor

import (
	"container/list"

	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/color/bytecolor"
	"github.com/asukakenji/drawing-challenge/common"
)

// Buffer is a canvas based on a buffer of bytecolor.Color.
// It implements the canvas.BufferBasedCanvas interface.
type Buffer struct {
	width           int
	height          int
	backgroundColor bytecolor.Color
	foregroundColor bytecolor.Color
	pixels          []bytecolor.Color
}

// Ensure that Buffer implements the canvas.BufferBasedCanvas interface.
var (
	_ canvas.BufferBasedCanvas = &Buffer{}
)

// NewBuffer returns a new Buffer.
//
// Errors
//
// common.ErrWidthOrHeightNotPositive:
// Will be returned if width <= 0, or height <= 0.
//
func NewBuffer(width, height int, bgColor, fgColor bytecolor.Color) (*Buffer, error) {
	if width <= 0 || height <= 0 {
		return nil, common.ErrWidthOrHeightNotPositive
	}
	pixels := make([]bytecolor.Color, width*height)
	fill(pixels, bgColor)
	return &Buffer{
		width:           width,
		height:          height,
		backgroundColor: bgColor,
		foregroundColor: fgColor,
		pixels:          pixels,
	}, nil
}

// Dimensions returns the width and height.
func (cnv *Buffer) Dimensions() (width, height int) {
	return cnv.width, cnv.height
}

// Pixels returns the underlying pixel buffer.
func (cnv *Buffer) Pixels() []bytecolor.Color {
	return cnv.pixels
}

// at is the same as At, but without boundary checks.
func (cnv *Buffer) at(x, y int) bytecolor.Color {
	index := xyToIndex(cnv.width, x, y)
	return cnv.pixels[index]
}

// At returns the color of the pixel at (x, y).
//
// Errors
//
// common.ErrPointOutsideCanvas:
// Will be returned if (x, y) is outside the canvas.
//
func (cnv *Buffer) At(x, y int) (color.Color, error) {
	if !isPointInsideCanvas(cnv.width, cnv.height, x, y) {
		return cnv.backgroundColor, common.ErrPointOutsideCanvas
	}
	return cnv.at(x, y), nil
}

// set is the same as Set, but without boundary checks.
func (cnv *Buffer) set(x, y int, bc bytecolor.Color) {
	index := xyToIndex(cnv.width, x, y)
	cnv.pixels[index] = bc
}

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
func (cnv *Buffer) Set(x, y int, c color.Color) error {
	if !isPointInsideCanvas(cnv.width, cnv.height, x, y) {
		return common.ErrPointOutsideCanvas
	}
	bc, ok := c.(bytecolor.Color)
	if !ok {
		return common.ErrColorTypeNotSupported
	}
	cnv.set(x, y, bc)
	return nil
}

// drawLine is the same as DrawLine, but without boundary checks.
func (cnv *Buffer) drawLine(x1, y1, x2, y2 int) {
	bc := cnv.foregroundColor
	if x1 == x2 {
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			cnv.set(x1, y, bc)
		}
	} else {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			cnv.set(x, y1, bc)
		}
	}
}

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
func (cnv *Buffer) DrawLine(x1, y1, x2, y2 int) error {
	if !isPointInsideCanvas(cnv.width, cnv.height, x1, y1) || !isPointInsideCanvas(cnv.width, cnv.height, x2, y2) {
		return common.ErrPointOutsideCanvas
	}
	// Check whether (x1, y1) and (x2, y2) are horizontally or vertically aligned
	if x1 != x2 && y1 != y2 {
		return common.ErrLineNotHorizontalOrVertical
	}
	cnv.drawLine(x1, y1, x2, y2)
	return nil
}

// DrawRect draws a rectangle.
//
// Errors
//
// common.ErrPointOutsideCanvas:
// Will be returned if (x1, y1) or (x2, y2) is outside the canvas.
//
func (cnv *Buffer) DrawRect(x1, y1, x2, y2 int) error {
	if !isPointInsideCanvas(cnv.width, cnv.height, x1, y1) || !isPointInsideCanvas(cnv.width, cnv.height, x2, y2) {
		return common.ErrPointOutsideCanvas
	}
	cnv.drawLine(x1, y1, x2, y1)
	cnv.drawLine(x1, y2, x2, y2)
	cnv.drawLine(x1, y1, x1, y2)
	cnv.drawLine(x2, y1, x2, y2)
	return nil
}

// bucketFill is the same as BucketFill, but without boundary checks.
func (cnv *Buffer) bucketFill(bc, colorToBeReplaced bytecolor.Color, pointsToBeFilled *list.List, pointsAlreadyProcessed *boolBuffer) {
	for pointsToBeFilled.Len() != 0 {
		back := pointsToBeFilled.Back()
		pointsToBeFilled.Remove(back)
		p := back.Value.(point)
		x, y := p.x, p.y
		c, err := cnv.At(x, y)
		if err != nil {
			continue
		}
		pointsAlreadyProcessed.Set(x, y)
		if !c.Equals(colorToBeReplaced) {
			continue
		}
		cnv.set(x, y, bc)
		if !pointsAlreadyProcessed.At(x-1, y) {
			pointsToBeFilled.PushBack(point{x - 1, y})
		}
		if !pointsAlreadyProcessed.At(x+1, y) {
			pointsToBeFilled.PushBack(point{x + 1, y})
		}
		if !pointsAlreadyProcessed.At(x, y-1) {
			pointsToBeFilled.PushBack(point{x, y - 1})
		}
		if !pointsAlreadyProcessed.At(x, y+1) {
			pointsToBeFilled.PushBack(point{x, y + 1})
		}
	}
}

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
func (cnv *Buffer) BucketFill(x, y int, c color.Color) error {
	if !isPointInsideCanvas(cnv.width, cnv.height, x, y) {
		return common.ErrPointOutsideCanvas
	}
	bc, ok := c.(bytecolor.Color)
	if !ok {
		return common.ErrColorTypeNotSupported
	}
	colorToBeReplaced := cnv.at(x, y)
	pointsToBeFilled := list.New()
	pointsToBeFilled.PushBack(point{x, y})
	pointsAlreadyProcessed := newBoolBuffer(cnv.width, cnv.height)
	cnv.bucketFill(bc, colorToBeReplaced, pointsToBeFilled, pointsAlreadyProcessed)
	return nil
}
