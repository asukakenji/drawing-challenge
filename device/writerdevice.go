package device

import (
	"fmt"
	"io"

	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

// WriterDevice TODO
type WriterDevice struct {
	writer io.Writer
}

// Ensure that riterDevice implements the Device interface
var (
	_ Device = &WriterDevice{}
)

// NewWriterDevice TODO
func NewWriterDevice(writer io.Writer) (*WriterDevice, error) {
	return &WriterDevice{
		writer: writer,
	}, nil
}

// renderTopBottomBorder TODO
func (dev *WriterDevice) renderTopBottomBorder(width int) error {
	// NOTE: Didn't use (width + 2) to prevent potential overflow
	fmt.Fprint(dev.writer, "-")
	for i := 0; i < width; i++ {
		fmt.Fprint(dev.writer, "-")
	}
	fmt.Fprintln(dev.writer, "-")
	return nil
}

// Render TODO
func (dev *WriterDevice) Render(cnv canvas.Canvas) error {
	bbcnv, ok := cnv.(canvas.BufferBasedCanvas)
	if !ok {
		return common.ErrCanvasNotSupported
	}
	width, height := bbcnv.Dimensions()
	dev.renderTopBottomBorder(width)
	if bcbcnv, ok := bbcnv.(*canvas.ByteColorBuffer); ok {
		pixels := bcbcnv.Pixels()
		offset := 0
		for j := 0; j < height; j++ {
			fmt.Fprintf(dev.writer, "|%s|\n", pixels[offset:offset+width])
			offset += width
		}
	} else {
		type toByter interface {
			ToByte() byte
		}
		for j := 0; j < height; j++ {
			fmt.Fprint(dev.writer, '|')
			for i := 0; i < width; i++ {
				c, err := bbcnv.At(i, j)
				if err != nil {
					// NOTE: This should not happen
					return err
				}
				if c2, ok := c.(color.ByteColor); ok {
					fmt.Fprint(dev.writer, c2)
				} else if c3, ok := c.(toByter); ok {
					fmt.Fprint(dev.writer, c3.ToByte())
				} else {
					return common.ErrInvalidColor
				}
			}
			fmt.Fprintln(dev.writer, '|')
		}
	}
	dev.renderTopBottomBorder(width)
	fmt.Fprintln(dev.writer)
	return nil
}
