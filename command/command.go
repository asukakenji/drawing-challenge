package command

import "github.com/asukakenji/drawing-challenge/color"

// Command represents a command defined in the project.
// There is no requirement on the underlying implementation.
// Types implementing Command are supposed to be value types
// with all their fields being public.
type Command interface {
	// Command is a dummy method to mark the type as implementing the Command interface.
	// See https://golang.org/src/go/ast/ast.go for an example in the standard library.
	Command()
}

// NewCanvasCommand represents the "new canvas" command.
type NewCanvasCommand struct {
	Width  int
	Height int
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd NewCanvasCommand) Command() {}

// DrawLineCommand represents the "draw line" command.
type DrawLineCommand struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd DrawLineCommand) Command() {}

// DrawRectCommand represents the "draw rect" command.
type DrawRectCommand struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd DrawRectCommand) Command() {}

// BucketFillCommand represents the "bucket fill" command.
type BucketFillCommand struct {
	X int
	Y int
	C color.Color
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd BucketFillCommand) Command() {}

// QuitCommand represents the "quit" command.
type QuitCommand struct {
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd QuitCommand) Command() {}

// Ensure that the command types implement the Command interface.
var (
	_ Command = NewCanvasCommand{}
	_ Command = DrawLineCommand{}
	_ Command = DrawRectCommand{}
	_ Command = BucketFillCommand{}
	_ Command = QuitCommand{}
)
