package command

import (
	"reflect"
	"testing"

	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

func TestNewBasicParser(t *testing.T) {
	colorParser := &color.ByteColorParser{
		DefaultColor: color.ByteColor(' '),
	}

	_, err := NewBasicParser(colorParser.ParseColor)
	if err != nil {
		t.Errorf("Expected: err == nil, Got: %#v", err)
	}

	_, err = NewBasicParser(nil)
	if err != common.ErrNilPointer {
		t.Errorf("Expected: err == %#v, Got: %#v", common.ErrNilPointer, err)
	}
}

func TestBasicParser_ParseCommand(t *testing.T) {
	colorParser := &color.ByteColorParser{
		DefaultColor: color.ByteColor(' '),
	}
	commandParser, err := NewBasicParser(colorParser.ParseColor)
	if err != nil {
		panic(err)
	}

	// Positive Cases
	casesPos := []struct {
		s       string
		command Command
	}{
		{"", EmptyCommand{}},
		{"C 20 4", NewCanvasCommand{20, 4}},                          // Example 1
		{"L 1 2 6 2", DrawLineCommand{1, 2, 6, 2}},                   // Example 2
		{"L 6 3 6 4", DrawLineCommand{6, 3, 6, 4}},                   // Example 3
		{"R 14 1 18 3", DrawRectCommand{14, 1, 18, 3}},               // Example 4
		{"B 10 3 o", BucketFillCommand{10, 3, color.ByteColor('o')}}, // Example 5
		{"Q", QuitCommand{}},                                         // Example 6
	}
	for _, c := range casesPos {
		command, err := commandParser.ParseCommand(c.s)
		if err != nil {
			t.Errorf("Case: %s, Expected: err == nil, Got: %#v", c.s, err)
		}
		if !reflect.DeepEqual(command, c.command) {
			t.Errorf("Case: %s, Expected: %#v, Got: %#v", c.s, c.command, command)
		}
	}

	// Negative Cases
	casesNeg := []struct {
		s   string
		err error
	}{
		{"C 1 2 3", common.ErrInvalidArgumentCount},
		{"C a 2", common.ErrInvalidNumber},
		{"C 1 b", common.ErrInvalidNumber},
		{"L 1 2 3 4 5", common.ErrInvalidArgumentCount},
		{"L a 2 3 4", common.ErrInvalidNumber},
		{"L 1 b 3 4", common.ErrInvalidNumber},
		{"L 1 2 c 4", common.ErrInvalidNumber},
		{"L 1 2 3 d", common.ErrInvalidNumber},
		{"R 1 2 3 4 5", common.ErrInvalidArgumentCount},
		{"R a 2 3 4", common.ErrInvalidNumber},
		{"R 1 b 3 4", common.ErrInvalidNumber},
		{"R 1 2 c 4", common.ErrInvalidNumber},
		{"R 1 2 3 d", common.ErrInvalidNumber},
		{"B 1 2 3 4", common.ErrInvalidArgumentCount},
		{"B a 2 o", common.ErrInvalidNumber},
		{"B 1 b o", common.ErrInvalidNumber},
		{"B 1 2 oo", common.ErrInvalidColor},
		{"X 20 4", common.ErrUnknownCommand},
	}
	for _, c := range casesNeg {
		_, err := commandParser.ParseCommand(c.s)
		if err != c.err {
			t.Errorf("Case: %s, Expected: %#v, Got: %#v", c.s, c.err, err)
		}
	}
}
