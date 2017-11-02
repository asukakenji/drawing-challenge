package canvas

import (
	"container/list"

	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

// ByteColorBuffer is a canvas based on a buffer of color.ByteColor.
// It implements the Canvas interface and the BufferBasedCanvas interface.
type ByteColorBuffer struct {
	width           int
	height          int
	backgroundColor color.ByteColor
	foregroundColor color.ByteColor
	pixels          []color.ByteColor
}

// Ensure that ByteColorBuffer implements the Canvas interface
// and the BufferBasedCanvas interface.
var (
	_ Canvas            = &ByteColorBuffer{}
	_ BufferBasedCanvas = &ByteColorBuffer{}
)

// NewByteColorBuffer returns a new ByteColorBuffer.
func NewByteColorBuffer(width, height int, bgColor, fgColor color.ByteColor) (*ByteColorBuffer, error) {
	if width <= 0 || height <= 0 {
		return nil, common.ErrWidthOrHeightNotPositive
	}
	pixels := make([]color.ByteColor, width*height)
	// TODO: Use the fast algorithm
	for i := range pixels {
		pixels[i] = bgColor
	}
	return &ByteColorBuffer{
		width:           width,
		height:          height,
		backgroundColor: bgColor,
		foregroundColor: fgColor,
		pixels:          pixels,
	}, nil
}

// Dimensions returns the width and height.
func (cnv *ByteColorBuffer) Dimensions() (width, height int) {
	return cnv.width, cnv.height
}

// Pixels returns the underlying pixel buffer.
func (cnv *ByteColorBuffer) Pixels() []color.ByteColor {
	return cnv.pixels
}

// set is the same as Set, but without boundary checks.
func (cnv *ByteColorBuffer) set(x, y int, bc color.ByteColor) {
	index := xyToIndex(cnv.width, x, y)
	cnv.pixels[index] = bc
}

// Set sets the color of the pixel at (x, y).
func (cnv *ByteColorBuffer) Set(x, y int, c color.Color) error {
	if !isPointInsideCanvas(cnv.width, cnv.height, x, y) {
		return common.ErrPointOutsideCanvas
	}
	bc, ok := c.(color.ByteColor)
	if !ok {
		return common.ErrColorTypeNotSupported
	}
	cnv.set(x, y, bc)
	return nil
}

// at is the same as At, but without boundary checks.
func (cnv *ByteColorBuffer) at(x, y int) color.ByteColor {
	index := xyToIndex(cnv.width, x, y)
	return cnv.pixels[index]
}

// At returns the color of the pixel at (x, y).
func (cnv *ByteColorBuffer) At(x, y int) (color.Color, error) {
	if !isPointInsideCanvas(cnv.width, cnv.height, x, y) {
		return cnv.backgroundColor, common.ErrPointOutsideCanvas
	}
	return cnv.at(x, y), nil
}

// drawLine is the same as DrawLine, but without boundary checks.
func (cnv *ByteColorBuffer) drawLine(x1, y1, x2, y2 int) {
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
func (cnv *ByteColorBuffer) DrawLine(x1, y1, x2, y2 int) error {
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
func (cnv *ByteColorBuffer) DrawRect(x1, y1, x2, y2 int) error {
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
func (cnv *ByteColorBuffer) bucketFill(bc color.ByteColor, pointsToBeFilled *list.List, pointsAlreadyProcessed *boolBuffer) {
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
		if c.Equals(cnv.foregroundColor) {
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

// BucketFill fills the area enclosing (x, y).
func (cnv *ByteColorBuffer) BucketFill(x, y int, c color.Color) error {
	if !isPointInsideCanvas(cnv.width, cnv.height, x, y) {
		return common.ErrPointOutsideCanvas
	}
	bc, ok := c.(color.ByteColor)
	if !ok {
		return common.ErrColorTypeNotSupported
	}
	pointsToBeFilled := list.New()
	pointsToBeFilled.PushBack(point{x, y})
	pointsAlreadyProcessed := newBoolBuffer(cnv.width, cnv.height)
	cnv.bucketFill(bc, pointsToBeFilled, pointsAlreadyProcessed)
	return nil
}
