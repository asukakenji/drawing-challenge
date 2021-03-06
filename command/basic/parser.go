package basic

import (
	"strconv"
	"strings"

	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/command"
	"github.com/asukakenji/drawing-challenge/common"
)

// Parser is a basic command parser.
// It implements the command.Parser interface.
//
// Commands supported by this parser:
// EmptyCommand,
// NewCanvasCommand,
// DrawLineCommand,
// DrawRectCommand,
// BucketFillCommand,
// QuitCommand.
//
type Parser struct {
	parseColorFunc func(string) (color.Color, error)
}

// Ensure that Parser implements the command.Parser interface.
var (
	_ command.Parser = &Parser{}
)

// NewParser returns a new Parser.
//
// Errors
//
// common.ErrNilPointer:
// Will be returned if parseColorFunc == nil.
//
func NewParser(parseColorFunc func(string) (color.Color, error)) (*Parser, error) {
	if parseColorFunc == nil {
		return nil, common.ErrNilPointer
	}
	return &Parser{
		parseColorFunc: parseColorFunc,
	}, nil
}

// ParseCommand parses the string s and returns a command.Command.
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
// Other errors
//
// common.ErrInvalidNumber:
// Will be returned when a numeric argument is expected,
// but it could not be parsed as a valid number.
//
// common.ErrInvalidColor:
// Will be returned when a color argument is expected,
// but it could not be parsed as a valid color.
//
func (parser *Parser) ParseCommand(s string) (command.Command, error) {
	if s == "" {
		return EmptyCommand{}, nil
	}
	words := strings.Split(s, " ")
	switch command, args := words[0], words[1:]; command {
	case "C":
		if len(args) != 2 {
			return nil, common.ErrInvalidArgumentCount
		}
		w, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		h, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		return NewCanvasCommand{w, h}, nil
	case "L":
		if len(args) != 4 {
			return nil, common.ErrInvalidArgumentCount
		}
		x1, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		y1, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		x2, err := strconv.Atoi(args[2])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		y2, err := strconv.Atoi(args[3])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		return DrawLineCommand{x1, y1, x2, y2}, nil
	case "R":
		if len(args) != 4 {
			return nil, common.ErrInvalidArgumentCount
		}
		x1, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		y1, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		x2, err := strconv.Atoi(args[2])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		y2, err := strconv.Atoi(args[3])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		return DrawRectCommand{x1, y1, x2, y2}, nil
	case "B":
		switch len(args) {
		case 2, 3:
			// OK
		default:
			return nil, common.ErrInvalidArgumentCount
		}
		x, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		y, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, common.ErrInvalidNumber
		}
		var colorString string
		if len(args) == 3 {
			colorString = args[2]
		}
		c, err := parser.parseColorFunc(colorString)
		if err != nil {
			return nil, err
		}
		return BucketFillCommand{x, y, c}, nil
	case "Q":
		return QuitCommand{}, nil
	default:
		return nil, common.ErrUnknownCommand
	}
}
