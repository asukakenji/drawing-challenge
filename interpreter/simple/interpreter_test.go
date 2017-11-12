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
type mockCanvas struct {
	width    int
	height   int
	commands *list.List
}

func newMockCanvas(width, height int) (canvas.Canvas, error) {
	if width <= 0 || height <= 0 {
		return nil, common.ErrWidthOrHeightNotPositive
	}
	commands := list.New()
	commands.PushBack(basic.NewCanvasCommand{Width: width, Height: height})
	return &mockCanvas{
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

// This type is created for testing purpose only
type mockCanvasContainer struct {
	cnv           canvas.Canvas
	newCanvasFunc func(int, int) (canvas.Canvas, error)
}

func (cc *mockCanvasContainer) Canvas() canvas.Canvas {
	return cc.cnv
}

func (cc *mockCanvasContainer) NewCanvas(width, height int) error {
	cnv, err := cc.newCanvasFunc(width, height)
	if err != nil {
		return err
	}
	cc.cnv = cnv
	return nil
}

// This type is created for testing purpose only
type mockRenderer struct{}

func (rdr *mockRenderer) Render(cnv canvas.Canvas) error {
	return nil
}

// This type is created for testing purpose only
type mockQuitter struct{}

func (qt *mockQuitter) ShouldQuit() bool {
	return false
}

func (qt *mockQuitter) SetQuit() {
}

// This type is created for testing purpose only
type mockEnvironment struct {
	mockCanvasContainer
	mockRenderer
	mockQuitter
}

func newMockEnvironment(newCanvasFunc func(int, int) (canvas.Canvas, error)) *mockEnvironment {
	return &mockEnvironment{
		mockCanvasContainer: mockCanvasContainer{
			newCanvasFunc: newCanvasFunc,
		},
	}
}

// This type is created for testing purpose only
type mockEnvironment1 struct {
	mockCanvasContainer
}

func newMockEnvironment1(newCanvasFunc func(int, int) (canvas.Canvas, error)) *mockEnvironment1 {
	return &mockEnvironment1{
		mockCanvasContainer: mockCanvasContainer{
			newCanvasFunc: newCanvasFunc,
		},
	}
}

// This type is created for testing purpose only
type mockEnvironment2 struct {
	mockCanvasContainer
	mockRenderer
}

func newMockEnvironment2(newCanvasFunc func(int, int) (canvas.Canvas, error)) *mockEnvironment2 {
	return &mockEnvironment2{
		mockCanvasContainer: mockCanvasContainer{
			newCanvasFunc: newCanvasFunc,
		},
	}
}

// This type is created for testing purpose only
type mockCommand struct{}

func (cmd mockCommand) Command() {
}

func TestNewInterpreter(t *testing.T) {
	_, err := NewInterpreter()
	if err != nil {
		t.Errorf("Expected: err == nil, Got: %#v", err)
	}
}

func TestInterpreter_Interpret(t *testing.T) {
	interp, err := NewInterpreter()
	if err != nil {
		panic(err)
	}

	// Positive Cases
	envPos := newMockEnvironment(newMockCanvas)

	casesPos := []struct {
		cmd    command.Command
		result command.Command
	}{
		{basic.NewCanvasCommand{Width: 20, Height: 4}, basic.NewCanvasCommand{Width: 20, Height: 4}},                                  // Example 1
		{basic.DrawLineCommand{X1: 1, Y1: 2, X2: 6, Y2: 2}, basic.DrawLineCommand{X1: 0, Y1: 1, X2: 5, Y2: 1}},                        // Example 2
		{basic.DrawLineCommand{X1: 6, Y1: 3, X2: 6, Y2: 4}, basic.DrawLineCommand{X1: 5, Y1: 2, X2: 5, Y2: 3}},                        // Example 3
		{basic.DrawRectCommand{X1: 14, Y1: 1, X2: 18, Y2: 3}, basic.DrawRectCommand{X1: 13, Y1: 0, X2: 17, Y2: 2}},                    // Example 4
		{basic.BucketFillCommand{X: 10, Y: 3, C: bytecolor.Color('o')}, basic.BucketFillCommand{X: 9, Y: 2, C: bytecolor.Color('o')}}, // Example 5
		{basic.EmptyCommand{}, basic.BucketFillCommand{X: 9, Y: 2, C: bytecolor.Color('o')}},
		{basic.QuitCommand{}, basic.BucketFillCommand{X: 9, Y: 2, C: bytecolor.Color('o')}},
	}
	for _, c := range casesPos {
		err = interp.Interpret(envPos, c.cmd)
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
	newCanvasFunc := func(width, height int) (canvas.Canvas, error) {
		return bc.NewBuffer(width, height, bytecolor.Color(' '), bytecolor.Color('x'))
	}

	envNeg := newMockEnvironment(newCanvasFunc)

	casesNeg := []struct {
		cmd command.Command
		err error
	}{
		// No Canvas
		{mockCommand{}, common.ErrCommandNotSupported},
		{basic.NewCanvasCommand{Width: -1, Height: -1}, common.ErrWidthOrHeightNotPositive},
		{basic.DrawLineCommand{X1: 1, Y1: 2, X2: 6, Y2: 2}, common.ErrCanvasNotCreated},
		{basic.DrawLineCommand{X1: 6, Y1: 3, X2: 6, Y2: 4}, common.ErrCanvasNotCreated},
		{basic.DrawRectCommand{X1: 14, Y1: 1, X2: 18, Y2: 3}, common.ErrCanvasNotCreated},
		{basic.BucketFillCommand{X: 10, Y: 3, C: bytecolor.Color('o')}, common.ErrCanvasNotCreated},
		// With Canvas
		{basic.NewCanvasCommand{Width: 20, Height: 4}, nil},
		{basic.DrawLineCommand{X1: -1, Y1: -1, X2: -1, Y2: -1}, common.ErrPointOutsideCanvas},
		{basic.DrawRectCommand{X1: -1, Y1: -1, X2: -1, Y2: -1}, common.ErrPointOutsideCanvas},
		{basic.BucketFillCommand{X: -1, Y: -1, C: bytecolor.Color('o')}, common.ErrPointOutsideCanvas},
	}
	for _, c := range casesNeg {
		err = interp.Interpret(envNeg, c.cmd)
		if err != c.err {
			t.Errorf("Case: (%#v, %#v), Expected: %#v, Got: %#v", envNeg, c.cmd, c.err, err)
		}
	}

	casesNeg2 := []struct {
		env interface{}
	}{
		{0},
		{newMockEnvironment1(newCanvasFunc)},
		{newMockEnvironment2(newCanvasFunc)},
	}
	for i, c := range casesNeg2 {
		err = interp.Interpret(c.env, basic.EmptyCommand{})
		if err != common.ErrEnvironmentNotSupported {
			t.Errorf("Case #%d: Expected: %#v, Got: %#v", i, common.ErrEnvironmentNotSupported, err)
		}
	}
}
