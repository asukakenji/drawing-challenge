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
