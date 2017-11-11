package simple

import (
	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/command"
	"github.com/asukakenji/drawing-challenge/common"
	"github.com/asukakenji/drawing-challenge/interpreter"
	"github.com/asukakenji/drawing-challenge/renderer"
)

// Interpreter is a simple command interpreter.
// It implements the interpreter.Interpreter interface.
//
// Commands supported by this interpreter:
// NewCanvasCommand,
// DrawLineCommand,
// DrawRectCommand,
// BucketFillCommand.
//
type Interpreter struct {
	newCanvasFunc func(int, int) (canvas.Canvas, error)
}

// Ensure that Interpreter implements the interpreter.Interpreter interface.
var (
	_ interpreter.Interpreter = &Interpreter{}
)

// NewInterpreter returns a new Interpreter.
//
// Errors
//
// common.ErrNilPointer:
// Will be returned if newCanvasFunc == nil.
//
func NewInterpreter(newCanvasFunc func(int, int) (canvas.Canvas, error)) (*Interpreter, error) {
	if newCanvasFunc == nil {
		return nil, common.ErrNilPointer
	}
	return &Interpreter{
		newCanvasFunc: newCanvasFunc,
	}, nil
}

// Interpret interprets the command cmd with the given environment env.
//
// env must implement the CanvasContainer interface
// and the renderer.Renderer interface.
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
	switch cmd := cmd.(type) {
	case command.NewCanvasCommand:
		cnv, err := interp.newCanvasFunc(cmd.Width, cmd.Height)
		if err != nil {
			return err
		}
		cc.SetCanvas(cnv)
		rdr.Render(cnv)
	case command.DrawLineCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.DrawLine(cmd.X1-1, cmd.Y1-1, cmd.X2-1, cmd.Y2-1)
		if err != nil {
			return err
		}
		rdr.Render(cnv)
	case command.DrawRectCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.DrawRect(cmd.X1-1, cmd.Y1-1, cmd.X2-1, cmd.Y2-1)
		if err != nil {
			return err
		}
		rdr.Render(cnv)
	case command.BucketFillCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.BucketFill(cmd.X-1, cmd.Y-1, cmd.C)
		if err != nil {
			return err
		}
		rdr.Render(cnv)
	default:
		return common.ErrCommandNotSupported
	}
	return nil
}
