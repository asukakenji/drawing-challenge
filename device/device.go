package device

import "github.com/asukakenji/drawing-challenge/canvas"

// Device represents a canvas renderer.
type Device interface {
	// Render renders cnv.
	// Implementations should specify the requirements on cnv.
	//
	// Errors
	//
	// common.ErrCanvasNotSupported:
	// Will be returned if cnv is not supported by this device.
	//
	// common.ErrColorNotSupported:
	// Will be returned if a color inside cnv is not supported by this device.
	//
	Render(cnv canvas.Canvas) error
}
