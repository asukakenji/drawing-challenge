package interpreter

import (
	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/command"
	"github.com/asukakenji/drawing-challenge/common"
)

// Interpreter represents a command interpreter.
type Interpreter interface {
	// Interpret interprets the command cmd with the given environment env.
	// Implementations should specify the requirements on env.
	//
	// Errors
	//
	// common.ErrEnvironmentNotSupported:
	// Will be returned if env is not supported by this interpreter.
	//
	// common.ErrCommandNotSupported:
	// Will be returned if cmd is not supported by this interpreter.
	//
	// Other errors:
	// May be returned depending on the commands supported.
	//
	Interpret(env interface{}, cmd command.Command) error
}

// BasicInterpreter is a basic command interpreter.
// It implements the Interpreter interface.
//
// Commands supported by this interpreter:
// NewCanvasCommand,
// DrawLineCommand,
// DrawRectCommand,
// BucketFillCommand.
//
type BasicInterpreter struct {
	newCanvasFunc func(int, int) (canvas.Canvas, error)
}

// Ensure that BasicInterpreter implements the Interpreter interface.
var (
	_ Interpreter = &BasicInterpreter{}
)

// NewBasicInterpreter returns a new BasicInterpreter.
//
// Errors
//
// common.ErrNilPointer:
// Will be returned if newCanvasFunc == nil.
//
func NewBasicInterpreter(newCanvasFunc func(int, int) (canvas.Canvas, error)) (*BasicInterpreter, error) {
	if newCanvasFunc == nil {
		return nil, common.ErrNilPointer
	}
	return &BasicInterpreter{
		newCanvasFunc: newCanvasFunc,
	}, nil
}

// CanvasContainer is a container of canvas.Canvas.
type CanvasContainer interface {
	// Canvas returns the contained canvas.Canvas.
	Canvas() canvas.Canvas

	// SetCanvas set the contained canvas.Canvas.
	SetCanvas(canvas.Canvas)
}

// Interpret interprets the command cmd with the given environment env.
//
// env must implement the CanvasContainer interface.
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
func (interp *BasicInterpreter) Interpret(env interface{}, cmd command.Command) error {
	cc, ok := env.(CanvasContainer)
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
	case command.DrawLineCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.DrawLine(cmd.X1-1, cmd.Y1-1, cmd.X2-1, cmd.Y2-1)
		if err != nil {
			return err
		}
	case command.DrawRectCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.DrawRect(cmd.X1-1, cmd.Y1-1, cmd.X2-1, cmd.Y2-1)
		if err != nil {
			return err
		}
	case command.BucketFillCommand:
		cnv := cc.Canvas()
		if cnv == nil {
			return common.ErrCanvasNotCreated
		}
		err := cnv.BucketFill(cmd.X-1, cmd.Y-1, cmd.C)
		if err != nil {
			return err
		}
	default:
		return common.ErrCommandNotSupported
	}
	return nil
}
