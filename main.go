package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/asukakenji/drawing-challenge/canvas"
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/command"
	"github.com/asukakenji/drawing-challenge/common"
	"github.com/asukakenji/drawing-challenge/device"
	"github.com/asukakenji/drawing-challenge/interpreter"
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

	// Setup background color
	bgColor, err := color.ParseByteColor(bgColorString)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup foreground color
	fgColor, err := color.ParseByteColor(fgColorString)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup environment
	env := &SimpleEnvironment{}

	// Setup device (standard output)
	dev, err := device.NewWriterDevice(os.Stdout)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup command parser
	parseColorFunc := func(s string) (color.Color, error) {
		return color.ParseByteColor(s)
	}
	commandParser, err := command.NewBasicParser(parseColorFunc)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup interpreter
	interp := &interpreter.BasicInterpreter{
		NewCanvasFunc: func(w, h int) (canvas.Canvas, error) {
			return canvas.NewByteColorBuffer(w, h, bgColor, fgColor)
		},
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
			if err == common.ErrEmptyCommand {
				continue
			}
			fmt.Println(err)
			continue
		}

		if _, ok := cmd.(command.QuitCommand); ok {
			return
		}

		err = interp.Interpret(env, cmd)
		if err != nil {
			fmt.Println(err)
		}
		if cnv := env.Canvas(); cnv != nil {
			dev.Render(cnv)
		}
	}
}
