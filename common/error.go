package common

import "errors"

var (
	// ErrNilPointer TODO
	ErrNilPointer = errors.New("Nil pointer")

	// ErrUnknownCommand TODO
	ErrUnknownCommand = errors.New("Unknown command")

	// ErrEmptyCommand TODO
	ErrEmptyCommand = errors.New("Empty command")

	// ErrCommandNotSupported TODO
	ErrCommandNotSupported = errors.New("Command not supported")

	// ErrInvalidArgumentCount TODO
	ErrInvalidArgumentCount = errors.New("Invalid number of arguments")

	// ErrInvalidNumber TODO
	ErrInvalidNumber = errors.New("Invalid number")

	// ErrInvalidColor TODO
	ErrInvalidColor = errors.New("Invalid color")

	// ErrEmptyColor TODO
	ErrEmptyColor = errors.New("Empty color")

	// ErrCanvasNotCreated TODO
	ErrCanvasNotCreated = errors.New("Canvas not created")

	// ErrCanvasNotSupported TODO
	ErrCanvasNotSupported = errors.New("Canvas not supported")
)
