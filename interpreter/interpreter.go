package interpreter

import "github.com/asukakenji/drawing-challenge/command"

// Interpreter represents a command interpreter.
type Interpreter interface {
	// Interpret interprets the command cmd with the given environment env.
	// Implementations should specify the requirements on env.
	Interpret(env interface{}, cmd command.Command) error
}
