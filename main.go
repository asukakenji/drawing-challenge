package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/command"
	"github.com/asukakenji/drawing-challenge/interpreter"
	"github.com/asukakenji/drawing-challenge/renderer"
)

// SimpleEnvironment TODO
type SimpleEnvironment struct {
	cnv canvas.Canvas
}

// Canvas TODO
func (env *SimpleEnvironment) Canvas() canvas.Canvas {
	return env.cnv
}

// SetCanvas TODO
func (env *SimpleEnvironment) SetCanvas(cnv canvas.Canvas) {
	env.cnv = cnv
}

const (
	// DefaultBGColorString TODO
	DefaultBGColorString = " "

	// DefaultFGColorString TODO
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
	colorParser := &color.ByteColorParser{
		DefaultColor: color.ByteColor(' '),
	}

	// Setup background color
	_bgColor, err := colorParser.ParseColor(bgColorString)
	if err != nil {
		fmt.Println(err)
		return
	}
	bgColor := _bgColor.(color.ByteColor)

	// Setup foreground color
	_fgColor, err := colorParser.ParseColor(fgColorString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fgColor := _fgColor.(color.ByteColor)

	// Setup environment
	env := &SimpleEnvironment{}

	// Setup device (standard output)
	dev, err := renderer.NewWriterRenderer(os.Stdout)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup command parser
	commandParser, err := command.NewBasicParser(colorParser.ParseColor)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup interpreter
	interp, err := interpreter.NewBasicInterpreter(func(width, height int) (canvas.Canvas, error) {
		return canvas.NewByteColorBuffer(width, height, bgColor, fgColor)
	})
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
		case command.EmptyCommand:
			continue
		case command.QuitCommand:
			return
		default:
			err = interp.Interpret(env, cmd)
			if err != nil {
				fmt.Println(err)
			}
			if cnv := env.Canvas(); cnv != nil {
				dev.Render(cnv)
			}
		}
	}
}
