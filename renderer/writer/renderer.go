// Package writer defines the Renderer type,
// which implements the renderer.Renderer interface.
package writer

import (
	"fmt"
	"io"

	"github.com/asukakenji/drawing-challenge/canvas"
	bc "github.com/asukakenji/drawing-challenge/canvas/bytecolor"
	"github.com/asukakenji/drawing-challenge/color/bytecolor"
	"github.com/asukakenji/drawing-challenge/common"
	"github.com/asukakenji/drawing-challenge/renderer"
)

// Renderer is a renderer based on an io.Writer.
// It implements the renderer.Renderer interface.
type Renderer struct {
	writer io.Writer
}

// Ensure that Renderer implements the renderer.Renderer interface.
var (
	_ renderer.Renderer = &Renderer{}
)

// NewRenderer returns a new Renderer.
//
// Errors
//
// common.ErrNilPointer:
// Will be returned if newCanvasFunc == nil.
//
func NewRenderer(writer io.Writer) (*Renderer, error) {
	if writer == nil {
		return nil, common.ErrNilPointer
	}
	return &Renderer{
		writer: writer,
	}, nil
}

// renderTopBottomBorder renders the top / buttom border of the canvas.
func (r *Renderer) renderTopBottomBorder(width int) {
	// NOTE: Didn't use (width + 2) to prevent potential overflow
	fmt.Fprint(r.writer, "-")
	for i := 0; i < width; i++ {
		fmt.Fprint(r.writer, "-")
	}
	fmt.Fprintln(r.writer, "-")
}

// Render renders cnv.
//
// cnv must implement the canvas.BufferBasedCanvas interface.
// The pixels of cnv must be color.ByteColor,
// or implements this method: "ToByte() byte".
//
// Errors
//
// common.ErrCanvasNotSupported:
// Will be returned if cnv is not supported by this renderer.
//
// common.ErrColorNotSupported:
// Will be returned if a color inside cnv is not supported by this renderer.
//
func (r *Renderer) Render(cnv canvas.Canvas) error {
	bbcnv, ok := cnv.(canvas.BufferBasedCanvas)
	if !ok {
		return common.ErrCanvasNotSupported
	}
	width, height := bbcnv.Dimensions()
	r.renderTopBottomBorder(width)
	if bcbcnv, ok := bbcnv.(*bc.Buffer); ok {
		// Render the canvas row-by-row if it is a ByteColorBuffer.
		pixels := bcbcnv.Pixels()
		offset := 0
		for j := 0; j < height; j++ {
			fmt.Fprintf(r.writer, "|%s|\n", pixels[offset:offset+width])
			offset += width
		}
	} else {
		// Render the canvas pixel-by-pixel if it is not a ByteColorBuffer.
		type toByter interface {
			ToByte() byte
		}
		for j := 0; j < height; j++ {
			fmt.Fprint(r.writer, '|')
			for i := 0; i < width; i++ {
				c, err := bbcnv.At(i, j)
				if err != nil {
					// NOTE: This should not happen if the canvas is correctly implemented
					panic(err)
				}
				if c2, ok := c.(bytecolor.Color); ok {
					fmt.Fprint(r.writer, c2)
				} else if c3, ok := c.(toByter); ok {
					fmt.Fprint(r.writer, c3.ToByte())
				} else {
					return common.ErrColorNotSupported
				}
			}
			fmt.Fprintln(r.writer, '|')
		}
	}
	r.renderTopBottomBorder(width)
	fmt.Fprintln(r.writer)
	return nil
}
