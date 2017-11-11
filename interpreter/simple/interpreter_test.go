package simple

import (
	"container/list"
	"reflect"
	"testing"

	"github.com/asukakenji/drawing-challenge/canvas"
	bc "github.com/asukakenji/drawing-challenge/canvas/bytecolor"
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/color/bytecolor"
	"github.com/asukakenji/drawing-challenge/command"
	"github.com/asukakenji/drawing-challenge/command/basic"
	"github.com/asukakenji/drawing-challenge/common"
)

// This type is created for testing purpose only
type mockCanvasContainer struct {
	cnv canvas.Canvas
}

func newMockCanvasContainer() *mockCanvasContainer {
	return &mockCanvasContainer{}
}

func (mcc *mockCanvasContainer) Canvas() canvas.Canvas {
	return mcc.cnv
}

func (mcc *mockCanvasContainer) SetCanvas(cnv canvas.Canvas) {
	mcc.cnv = cnv
}

// This type is created for testing purpose only
type mockCanvas struct {
	width    int
	height   int
	commands *list.List
}

func NewMockCanvas(width, height int) (canvas.Canvas, error) {
	commands := list.New()
	commands.PushBack(basic.NewCanvasCommand{Width: width, Height: height})
	return &mockCanvas{
		width:    width,
		height:   height,
		commands: commands,
	}, nil
}

func (mc *mockCanvas) Dimensions() (int, int) {
	return mc.width, mc.height
}

func (mc *mockCanvas) DrawLine(x1, y1, x2, y2 int) error {
	mc.commands.PushBack(basic.DrawLineCommand{X1: x1, Y1: y1, X2: x2, Y2: y2})
	return nil
}

func (mc *mockCanvas) DrawRect(x1, y1, x2, y2 int) error {
	mc.commands.PushBack(basic.DrawRectCommand{X1: x1, Y1: y1, X2: x2, Y2: y2})
	return nil
}

func (mc *mockCanvas) BucketFill(x, y int, c color.Color) error {
	mc.commands.PushBack(basic.BucketFillCommand{X: x, Y: y, C: c})
	return nil
}

func TestNewInterpreter(t *testing.T) {
	_, err := NewInterpreter(NewMockCanvas)
	if err != nil {
		t.Errorf("Expected: err == nil, Got: %#v", err)
	}

	_, err = NewInterpreter(nil)
	if err != common.ErrNilPointer {
		t.Errorf("Expected: err == %#v, Got: %#v", common.ErrNilPointer, err)
	}
}

func TestInterpreter_Interpret(t *testing.T) {
	// Positive Cases
	interpPos, err := NewInterpreter(NewMockCanvas)
	if err != nil {
		panic(err)
	}

	envPos := newMockCanvasContainer()

	casesPos := []struct {
		cmd    command.Command
		result command.Command
	}{
		{basic.NewCanvasCommand{Width: 20, Height: 4}, basic.NewCanvasCommand{Width: 20, Height: 4}},                                  // Example 1
		{basic.DrawLineCommand{X1: 1, Y1: 2, X2: 6, Y2: 2}, basic.DrawLineCommand{X1: 0, Y1: 1, X2: 5, Y2: 1}},                        // Example 2
		{basic.DrawLineCommand{X1: 6, Y1: 3, X2: 6, Y2: 4}, basic.DrawLineCommand{X1: 5, Y1: 2, X2: 5, Y2: 3}},                        // Example 3
		{basic.DrawRectCommand{X1: 14, Y1: 1, X2: 18, Y2: 3}, basic.DrawRectCommand{X1: 13, Y1: 0, X2: 17, Y2: 2}},                    // Example 4
		{basic.BucketFillCommand{X: 10, Y: 3, C: bytecolor.Color('o')}, basic.BucketFillCommand{X: 9, Y: 2, C: bytecolor.Color('o')}}, // Example 5
	}
	for _, c := range casesPos {
		err = interpPos.Interpret(envPos, c.cmd)
		if err != nil {
			t.Errorf("Case: (%#v, %#v), Expected: err == nil, Got: %#v", envPos, c.cmd, err)
		}
		mc := envPos.Canvas().(*mockCanvas)
		cmd := mc.commands.Back().Value
		if !reflect.DeepEqual(cmd, c.result) {
			t.Errorf("Case: (%#v, %#v), Expected: %#v, Got: %#v", envPos, c.cmd, c.result, cmd)
		}
	}

	// Negative Cases
	interpNeg, err := NewInterpreter(func(width, height int) (canvas.Canvas, error) {
		return bc.NewBuffer(width, height, bytecolor.Color(' '), bytecolor.Color('x'))
	})
	if err != nil {
		panic(err)
	}

	envNeg := newMockCanvasContainer()

	casesNeg := []struct {
		cmd command.Command
		err error
	}{
		// No Canvas
		{basic.EmptyCommand{}, common.ErrCommandNotSupported},
		{basic.NewCanvasCommand{Width: -1, Height: -1}, common.ErrWidthOrHeightNotPositive},
		{basic.DrawLineCommand{X1: 1, Y1: 2, X2: 6, Y2: 2}, common.ErrCanvasNotCreated},
		{basic.DrawLineCommand{X1: 6, Y1: 3, X2: 6, Y2: 4}, common.ErrCanvasNotCreated},
		{basic.DrawRectCommand{X1: 14, Y1: 1, X2: 18, Y2: 3}, common.ErrCanvasNotCreated},
		{basic.BucketFillCommand{X: 10, Y: 3, C: bytecolor.Color('o')}, common.ErrCanvasNotCreated},
		{basic.QuitCommand{}, common.ErrCommandNotSupported},
		// With Canvas
		{basic.NewCanvasCommand{Width: 20, Height: 4}, nil},
		{basic.DrawLineCommand{X1: -1, Y1: -1, X2: -1, Y2: -1}, common.ErrPointOutsideCanvas},
		{basic.DrawRectCommand{X1: -1, Y1: -1, X2: -1, Y2: -1}, common.ErrPointOutsideCanvas},
		{basic.BucketFillCommand{X: -1, Y: -1, C: bytecolor.Color('o')}, common.ErrPointOutsideCanvas},
	}
	for _, c := range casesNeg {
		err = interpNeg.Interpret(envNeg, c.cmd)
		if err != c.err {
			t.Errorf("Case: (%#v, %#v), Expected: %#v, Got: %#v", envNeg, c.cmd, c.err, err)
		}
	}

	err = interpNeg.Interpret(0, basic.EmptyCommand{})
	if err != common.ErrEnvironmentNotSupported {
		t.Errorf("Case: env == 0, Expected: %#v, Got: %#v", common.ErrEnvironmentNotSupported, err)
	}
}
