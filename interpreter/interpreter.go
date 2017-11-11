// Package interpreter defines the Interpreter interface.
package interpreter

import "github.com/asukakenji/drawing-challenge/command"

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
