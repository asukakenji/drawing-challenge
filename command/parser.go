package command

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
