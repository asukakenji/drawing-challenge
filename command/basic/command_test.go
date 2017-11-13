package basic

import (
	"testing"

	"github.com/asukakenji/drawing-challenge/command"
)

// This is created to fulfill Coveralls (https://coveralls.io/)
// Normally this is not needed.
func TestCommand(t *testing.T) {
	cases := []struct {
		cmd command.Command
	}{
		{EmptyCommand{}},
		{NewCanvasCommand{}},
		{DrawLineCommand{}},
		{DrawRectCommand{}},
		{BucketFillCommand{}},
		{QuitCommand{}},
	}
	for _, c := range cases {
		c.cmd.Command()
	}
}
