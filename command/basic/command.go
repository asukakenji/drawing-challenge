// Package basic defines several "Value Object" types
// which implement the command.Command interface,
// and the Parser type,
// which implements the command.Parser interface.
package basic

import (
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/command"
)

// EmptyCommand represents a "No-op" command.
// It implements the Command interface.
type EmptyCommand struct {
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd EmptyCommand) Command() {}

// NewCanvasCommand represents the "new canvas" command.
// It implements the Command interface.
type NewCanvasCommand struct {
	Width  int
	Height int
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd NewCanvasCommand) Command() {}

// DrawLineCommand represents the "draw line" command.
// It implements the Command interface.
type DrawLineCommand struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd DrawLineCommand) Command() {}

// DrawRectCommand represents the "draw rect" command.
// It implements the Command interface.
type DrawRectCommand struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd DrawRectCommand) Command() {}

// BucketFillCommand represents the "bucket fill" command.
// It implements the Command interface.
type BucketFillCommand struct {
	X int
	Y int
	C color.Color
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd BucketFillCommand) Command() {}

// QuitCommand represents the "quit" command.
// It implements the Command interface.
type QuitCommand struct {
}

// Command is a dummy method to mark the type as implementing the Command interface.
func (cmd QuitCommand) Command() {}

// Ensure that the command types implement the command.Command interface.
var (
	_ command.Command = EmptyCommand{}
	_ command.Command = NewCanvasCommand{}
	_ command.Command = DrawLineCommand{}
	_ command.Command = DrawRectCommand{}
	_ command.Command = BucketFillCommand{}
	_ command.Command = QuitCommand{}
)
