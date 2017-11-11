// Package renderer defines the Renderer interface.
package renderer

import "github.com/asukakenji/drawing-challenge/canvas"

// Renderer represents a canvas renderer.
type Renderer interface {
	// Render renders cnv.
	// Implementations should specify the requirements on cnv.
	//
	// Errors
	//
	// common.ErrCanvasNotSupported:
	// Will be returned if cnv is not supported by this renderer.
	//
	// common.ErrColorNotSupported:
	// Will be returned if a color inside cnv is not supported by this renderer.
	//
	Render(cnv canvas.Canvas) error
}
