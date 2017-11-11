package main

import (
	"bufio"
	"flag"
	"fmt"
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
		fmt.Println(err)
		return
	}
	bgColor := _bgColor.(bytecolor.Color)

	// Setup foreground color
	_fgColor, err := colorParser.ParseColor(fgColorString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fgColor := _fgColor.(bytecolor.Color)

	// Setup command parser
	commandParser, err := basic.NewParser(colorParser.ParseColor)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup interpreter
	interp, err := simple.NewInterpreter(func(width, height int) (canvas.Canvas, error) {
		return bc.NewBuffer(width, height, bgColor, fgColor)
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup renderer (standard output)
	rdr, err := writer.NewRenderer(os.Stdout)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup environment
	env, err := simple.NewEnvironment(rdr)
	if err != nil {
		fmt.Println(err)
		return
	}

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Split(bufio.ScanLines)
	for {
		fmt.Print("enter command: ")
		if !stdin.Scan() {
			break
		}
		line := stdin.Text()
		cmd, err := commandParser.ParseCommand(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch cmd := cmd.(type) {
		case basic.EmptyCommand:
			continue
		case basic.QuitCommand:
			return
		default:
			err = interp.Interpret(env, cmd)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
