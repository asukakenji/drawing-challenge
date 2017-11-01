package device

import "github.com/asukakenji/drawing-challenge/canvas"

// Device TODO
type Device interface {
	Render(cnv canvas.Canvas) error
}
