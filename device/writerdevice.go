package device

import (
	"fmt"
	"io"

	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

// WriterDevice is a Device based on a Writer.
type WriterDevice struct {
	writer io.Writer
}

// Ensure that WriterDevice implements the Device interface.
var (
	_ Device = &WriterDevice{}
)

// NewWriterDevice returns a new WriterDevice.
func NewWriterDevice(writer io.Writer) (*WriterDevice, error) {
	return &WriterDevice{
		writer: writer,
	}, nil
}

// renderTopBottomBorder renders the top / buttom border of the canvas.
func (dev *WriterDevice) renderTopBottomBorder(width int) error {
	// NOTE: Didn't use (width + 2) to prevent potential overflow
	fmt.Fprint(dev.writer, "-")
	for i := 0; i < width; i++ {
		fmt.Fprint(dev.writer, "-")
	}
	fmt.Fprintln(dev.writer, "-")
	return nil
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
// Will be returned if cnv is not supported by this device.
//
// common.ErrColorNotSupported:
// Will be returned if a color inside cnv is not supported by this device.
//
func (dev *WriterDevice) Render(cnv canvas.Canvas) error {
	bbcnv, ok := cnv.(canvas.BufferBasedCanvas)
	if !ok {
		return common.ErrCanvasNotSupported
	}
	width, height := bbcnv.Dimensions()
	dev.renderTopBottomBorder(width)
	if bcbcnv, ok := bbcnv.(*canvas.ByteColorBuffer); ok {
		// Render the canvas row-by-row if it is a ByteColorBuffer.
		pixels := bcbcnv.Pixels()
		offset := 0
		for j := 0; j < height; j++ {
			fmt.Fprintf(dev.writer, "|%s|\n", pixels[offset:offset+width])
			offset += width
		}
	} else {
		// Render the canvas pixel-by-pixel if it is not a ByteColorBuffer.
		type toByter interface {
			ToByte() byte
		}
		for j := 0; j < height; j++ {
			fmt.Fprint(dev.writer, '|')
			for i := 0; i < width; i++ {
				c, err := bbcnv.At(i, j)
				if err != nil {
					// NOTE: This should not happen
					panic(err)
				}
				if c2, ok := c.(color.ByteColor); ok {
					fmt.Fprint(dev.writer, c2)
				} else if c3, ok := c.(toByter); ok {
					fmt.Fprint(dev.writer, c3.ToByte())
				} else {
					return common.ErrColorNotSupported
				}
			}
			fmt.Fprintln(dev.writer, '|')
		}
	}
	dev.renderTopBottomBorder(width)
	fmt.Fprintln(dev.writer)
	return nil
}
