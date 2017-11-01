package interpreter

import "github.com/asukakenji/drawing-challenge/command"

// Interpreter TODO
type Interpreter interface {
	Interpret(env interface{}, cmd command.Command) error
}
