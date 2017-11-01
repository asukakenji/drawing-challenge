package interpreter

import (
	"fmt"

	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/command"
	"github.com/asukakenji/drawing-challenge/common"
)

// Ensure that BasicInterpreter implements the Interpreter interface
var (
	_ Interpreter = &BasicInterpreter{}
)

// BasicInterpreter TODO
type BasicInterpreter struct {
	NewCanvasFunc func(int, int) (canvas.Canvas, error)
}

// Interpret TODO
func (interp *BasicInterpreter) Interpret(env interface{}, cmd command.Command) error {
	type canvasContainer interface {
		Canvas() canvas.Canvas
		SetCanvas(canvas.Canvas)
	}
	cc, ok := env.(canvasContainer)
	if !ok {
		return fmt.Errorf("Environment not supported")
	}
	switch cmd := cmd.(type) {
	case command.NewCanvasCommand:
		cnv, err := interp.NewCanvasFunc(cmd.Width, cmd.Height)
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
