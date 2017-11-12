// Package simple defines the Interpreter type,
// which is a stateless interpreter implementing interpreter.Interpreter,
// and the CanvasContainer interface and the Quitter interface,
// which are used to specify the requirements of the Interpreter type,
// and the Environment type, which fulfills the requirements.
package simple

import (
	"github.com/asukakenji/drawing-challenge/command"
	"github.com/asukakenji/drawing-challenge/command/basic"
	"github.com/asukakenji/drawing-challenge/common"
	"github.com/asukakenji/drawing-challenge/interpreter"
	"github.com/asukakenji/drawing-challenge/renderer"
)

// Interpreter is a simple command interpreter.
// It implements the interpreter.Interpreter interface.
//
// Commands supported by this interpreter:
// basic.EmptyCommand,
// basic.NewCanvasCommand,
// basic.DrawLineCommand,
// basic.DrawRectCommand,
// basic.BucketFillCommand,
// basic.QuitCommand.
//
type Interpreter struct {
}

// Ensure that Interpreter implements the interpreter.Interpreter interface.
var (
	_ interpreter.Interpreter = &Interpreter{}
)

// NewInterpreter returns a new Interpreter.
//
// Errors
//
// (None)
//
func NewInterpreter() (*Interpreter, error) {
	return &Interpreter{}, nil
}

// Interpret interprets the command cmd with the given environment env.
//
// env must implement the CanvasContainer interface,
// the renderer.Renderer interface,
// and the Quitter interface.
//
// Errors
//
// common.ErrEnvironmentNotSupported:
// Will be returned if env is not supported by this interpreter.
//
// common.ErrCommandNotSupported:
// Will be returned if cmd is not supported by this interpreter.
//
// Other errors
//
// common.ErrCanvasNotCreated:
// Will be returned if a canvas is needed, but it has not been created.
//
// Errors returned from the newCanvasFunc function, the canvas' DrawLine,
// DrawRect, and BucketFill methods are returned without modifications.
//
func (interp *Interpreter) Interpret(env interface{}, cmd command.Command) error {
	cc, ok := env.(CanvasContainer)
	if !ok {
		return common.ErrEnvironmentNotSupported
	}
	rdr, ok := env.(renderer.Renderer)
	if !ok {
		return common.ErrEnvironmentNotSupported
	}
	qt, ok := env.(Quitter)
	if !ok {
		return common.ErrEnvironmentNotSupported
	}
	switch cmd := cmd.(type) {
	case basic.EmptyCommand:
		// Nothing to be done
	case basic.NewCanvasCommand:
		err := cc.NewCanvas(cmd.Width, cmd.Height)
		if err != nil {
			return err
		}
		cnv := cc.Canvas()
		rdr.Render(cnv)
	case basic.DrawLineCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.DrawLine(cmd.X1-1, cmd.Y1-1, cmd.X2-1, cmd.Y2-1)
		if err != nil {
			return err
		}
		rdr.Render(cnv)
	case basic.DrawRectCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.DrawRect(cmd.X1-1, cmd.Y1-1, cmd.X2-1, cmd.Y2-1)
		if err != nil {
			return err
		}
		rdr.Render(cnv)
	case basic.BucketFillCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.BucketFill(cmd.X-1, cmd.Y-1, cmd.C)
		if err != nil {
			return err
		}
		rdr.Render(cnv)
	case basic.QuitCommand:
		qt.SetQuit()
	default:
		return common.ErrCommandNotSupported
	}
	return nil
}
