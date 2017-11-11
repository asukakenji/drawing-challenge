// Package common defines types and variables
// which are needed by other packages in the project.
package common

import "errors"

var (
	// ErrNilPointer indicates the argument is nil where a non-nil value is needed.
	ErrNilPointer = errors.New("Nil pointer")

	// ---

	// ErrCanvasNotSupported indicates the canvas is not supported by the renderer.
	ErrCanvasNotSupported = errors.New("Canvas not supported")

	// ErrColorNotSupported indicates the color is not supported by the renderer.
	ErrColorNotSupported = errors.New("Color not supported")

	// ---

	// ErrEnvironmentNotSupported indicates the environment is not supported by the interpreter.
	ErrEnvironmentNotSupported = errors.New("Environment not supported")

	// ErrCommandNotSupported indicates the command is not supported by the interpreter.
	ErrCommandNotSupported = errors.New("Command not supported")

	// ErrCanvasNotCreated indicates the canvas is not created where a command needs it.
	ErrCanvasNotCreated = errors.New("Canvas not created")

	// ---

	// ErrUnknownCommand indicates the command is not recognized by the command parser.
	ErrUnknownCommand = errors.New("Unknown command")

	// ErrInvalidArgumentCount indicates the number of arguments of the command is invalid.
	ErrInvalidArgumentCount = errors.New("Invalid number of arguments")

	// ErrInvalidNumber indicates an argument could not be parsed to a number.
	ErrInvalidNumber = errors.New("Invalid number")

	// ---

	// ErrWidthOrHeightNotPositive indicates the width or height of the canvas is not positive.
	ErrWidthOrHeightNotPositive = errors.New("'width' or 'height' not positive")

	// ErrPointOutsideCanvas indicates the point specified is outside the canvas.
	ErrPointOutsideCanvas = errors.New("Point outside canvas")

	// ErrColorTypeNotSupported indicates the type of the color is not supported by the canvas.
	ErrColorTypeNotSupported = errors.New("Color type not supported")

	// ErrLineNotHorizontalOrVertical indicates the line specified is not horizontal or vertical.
	ErrLineNotHorizontalOrVertical = errors.New("Line not horizontal or vertical")

	// ---

	// ErrInvalidColor indicates the argument could not be parseed to a color value.
	ErrInvalidColor = errors.New("Invalid color")
)
