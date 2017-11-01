package device

import "github.com/asukakenji/drawing-challenge/canvas"

// Device is an entity that renders a Canvas.
type Device interface {
	// Render renders cnv.
	// Implementations should specify what kind of canvas they could render.
	Render(cnv canvas.Canvas) error
}
