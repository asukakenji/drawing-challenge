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

	// NewCanvas creates a new canvas.Canvas.
	NewCanvas(width, height int) error
}

// Quitter is a container of a bool which determines if the program should quit.
type Quitter interface {
	// ShouldQuit returns if the program should quit.
	ShouldQuit() bool

	// SetQuit causes ShouldQuit to return true hereafter.
	SetQuit()
}

// Environment is a simple environment for the interpreter.
// It implements the CanvasContainer interface
// and the renderer.Renderer interface.
type Environment struct {
	newCanvasFunc func(int, int) (canvas.Canvas, error)
	cnv           canvas.Canvas
	rdr           renderer.Renderer
	shouldQuit    bool
}

// NewEnvironment returns a new Environment.
//
// Errors
//
// common.ErrNilPointer:
// Will be returned if rdr == nil.
//
func NewEnvironment(newCanvasFunc func(int, int) (canvas.Canvas, error), rdr renderer.Renderer) (*Environment, error) {
	if newCanvasFunc == nil {
		return nil, common.ErrNilPointer
	}
	if rdr == nil {
		return nil, common.ErrNilPointer
	}
	return &Environment{
		newCanvasFunc: newCanvasFunc,
		rdr:           rdr,
	}, nil
}

// Canvas returns the contained canvas.Canvas.
func (env *Environment) Canvas() canvas.Canvas {
	return env.cnv
}

// NewCanvas creates a new canvas.Canvas.
func (env *Environment) NewCanvas(width, height int) error {
	cnv, err := env.newCanvasFunc(width, height)
	if err != nil {
		return err
	}
	env.cnv = cnv
	return nil
}

// Render renders cnv.
func (env *Environment) Render(cnv canvas.Canvas) error {
	return env.rdr.Render(cnv)
}

// ShouldQuit returns if the program should quit.
func (env *Environment) ShouldQuit() bool {
	return env.shouldQuit
}

// SetQuit causes ShouldQuit to return true hereafter.
func (env *Environment) SetQuit() {
	env.shouldQuit = true
}
