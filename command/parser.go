package command

import (
	"strconv"
	"strings"

	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

// Parser TODO
type Parser interface {
	ParseCommand(string) (Command, error)
}

// BasicParser TODO
type BasicParser struct {
	parseColorFunc func(string) (color.Color, error)
}

// Ensure that BasicParser implements the Parser interface
var (
	_ Parser = &BasicParser{}
)

// NewBasicParser TODO
func NewBasicParser(parseColorFunc func(string) (color.Color, error)) (*BasicParser, error) {
	if parseColorFunc == nil {
		return nil, common.ErrNilPointer
	}
	return &BasicParser{
		parseColorFunc: parseColorFunc,
	}, nil
}

// ParseCommand TODO
func (parser *BasicParser) ParseCommand(s string) (Command, error) {
	if s == "" {
		return nil, common.ErrEmptyCommand
	}
	// TODO: Use regular expression
	words := strings.Split(s, " ")
	// TODO: Check len(words)
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
