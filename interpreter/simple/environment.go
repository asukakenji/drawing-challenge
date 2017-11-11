// Package simple defines the CanvasContainer interface,
// and the Environment type
// which implements it and the renderer.Renderer interface.
package simple

import (
	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/common"
	"github.com/asukakenji/drawing-challenge/renderer"
)

// CanvasContainer is a container of canvas.Canvas.
type CanvasContainer interface {
	// Canvas returns the contained canvas.Canvas.
	Canvas() canvas.Canvas

	// SetCanvas set the contained canvas.Canvas.
	SetCanvas(canvas.Canvas)
}

// Environment is a simple environment for the interpreter.
// It implements the CanvasContainer interface
// and the renderer.Renderer interface.
type Environment struct {
	cnv canvas.Canvas
	rdr renderer.Renderer
}

// NewEnvironment returns a new Environment.
//
// Errors
//
// common.ErrNilPointer:
// Will be returned if rdr == nil.
//
func NewEnvironment(rdr renderer.Renderer) (*Environment, error) {
	if rdr == nil {
		return nil, common.ErrNilPointer
	}
	return &Environment{
		rdr: rdr,
	}, nil
}

// Canvas returns the contained canvas.Canvas.
func (env *Environment) Canvas() canvas.Canvas {
	return env.cnv
}

// SetCanvas set the contained canvas.Canvas.
func (env *Environment) SetCanvas(cnv canvas.Canvas) {
	env.cnv = cnv
}

// Render renders cnv.
func (env *Environment) Render(cnv canvas.Canvas) error {
	return env.rdr.Render(cnv)
}
