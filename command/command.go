// Package command defines the Command interface, and the Parser interface.
package command

// Command represents a command defined in the project.
//
// There is no requirement on the underlying implementation.
// Types implementing Command are supposed to be value types
// with all their fields being public.
type Command interface {
	// Command is a dummy method to mark the type as implementing the Command interface.
	Command()
}

// NOTE:
// See https://golang.org/src/go/ast/ast.go for examples of
// dummy interface methods in the standard library.

// Parser represents a command parser.
type Parser interface {
	// ParseCommand parses s and returns a Command.
	//
	// Errors
	//
	// common.ErrUnknownCommand:
	// Will be returned if s contains a command not recognized by this parser.
	//
	// common.ErrInvalidArgumentCount:
	// Will be returned if s contains a command recognized by this parser,
	// but the argument count is invalid.
	//
	// Other errors:
	// May be returned depending on the commands supported.
	//
	ParseCommand(s string) (Command, error)
}
