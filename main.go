package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/asukakenji/drawing-challenge/canvas"
	bc "github.com/asukakenji/drawing-challenge/canvas/bytecolor"
	"github.com/asukakenji/drawing-challenge/color/bytecolor"
	"github.com/asukakenji/drawing-challenge/command/basic"
	"github.com/asukakenji/drawing-challenge/interpreter/simple"
	"github.com/asukakenji/drawing-challenge/renderer/writer"
)

const (
	// DefaultBGColorString is the default value for bgColorString.
	DefaultBGColorString = " "

	// DefaultFGColorString is the default value for fgColorString.
	DefaultFGColorString = "x"
)

var (
	bgColorString string
	fgColorString string
)

func init() {
	flag.StringVar(&bgColorString, "bgColor", DefaultBGColorString, "The background color of the canvas")
	flag.StringVar(&fgColorString, "fgColor", DefaultFGColorString, "The foreground color of the canvas")
}

var (
	input  io.Reader = os.Stdin
	output io.Writer = os.Stdout
)

func main() {
	// Setup command line flags
	flag.Parse()

	// Setup color parser
	colorParser := &bytecolor.Parser{
		DefaultColor: bytecolor.Color(' '),
	}

	// Setup background color
	_bgColor, err := colorParser.ParseColor(bgColorString)
	if err != nil {
		panic(err)
	}
	bgColor := _bgColor.(bytecolor.Color)

	// Setup foreground color
	_fgColor, err := colorParser.ParseColor(fgColorString)
	if err != nil {
		panic(err)
	}
	fgColor := _fgColor.(bytecolor.Color)

	// Setup command parser (the only possible error is common.ErrNilPointer)
	commandParser, _ := basic.NewParser(colorParser.ParseColor)

	// Setup interpreter (no error)
	interp, _ := simple.NewInterpreter()

	// Setup renderer (the only possible error is common.ErrNilPointer)
	rdr, err := writer.NewRenderer(output)
	if err != nil {
		panic(err)
	}

	// Setup environment (the only possible error is common.ErrNilPointer)
	newCanvasFunc := func(width, height int) (canvas.Canvas, error) {
		return bc.NewBuffer(width, height, bgColor, fgColor)
	}
	env, _ := simple.NewEnvironment(newCanvasFunc, rdr)

	stdin := bufio.NewScanner(input)
	stdin.Split(bufio.ScanLines)
	for !env.ShouldQuit() {
		fmt.Fprint(output, "enter command: ")
		if !stdin.Scan() {
			break
		}
		line := stdin.Text()
		cmd, err := commandParser.ParseCommand(line)
		if err != nil {
			fmt.Fprintln(output, err)
			continue
		}

		err = interp.Interpret(env, cmd)
		if err != nil {
			fmt.Fprintln(output, err)
		}
	}
}
