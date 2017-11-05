# Drawing Challenge

The latest version of this project is available at:

https://github.com/asukakenji/drawing-challenge

## Manuals

- [User Manual](#user-manual)
- [Technical Diagrams](#technical-diagrams)
- [Design Documentation](#design-documentation)
- [API Documentation](#api-documentation)

## User Manual

### Installation (From GitHub, Preferred Way)

1. Download Go from the official web site [here](https://golang.org/dl/).
1. Follow the installation instructions [here](https://golang.org/doc/install) to complete the installation.
   - The most important point is setting the `PATH` and `GOPATH` environment variables correctly.
1. Start a command prompt / terminal
1. Enter the following command to download the source from GitHub:

        go get -u github.com/asukakenji/drawing-challenge
1. Enter the following commands to execute the program:

    UNIX-based operating systems:

        $GOPATH/bin/drawing-challenge

    Windows:

        %GOPATH%\bin\drawing-challenge.exe

### Installation (From Archive)

1. Download Go from the official web site [here](https://golang.org/dl/).
1. Follow the installation instructions [here](https://golang.org/doc/install) to complete the installation.
   - The most important point is setting the `PATH` and `GOPATH` environment variables correctly.
1. Download the source archive.
1. Decompress the source archive to the correct directory:

    UNIX-based operating systems:

        $GOPATH/src/github.com/asukakenji/drawing-challenge

    Windows:

        %GOPATH%\src\github.com\asukakenji\drawing-challenge
1. Start a command prompt / terminal
1. Enter the following commands to execute the program without compiling:

    UNIX-based operating systems:

        cd $GOPATH/src/github.com/asukakenji/drawing-challenge
        go run ./main.go

    Windows:

        cd %GOPATH%\src\github.com\asukakenji\drawing-challenge
        go run .\main.go
1. Or, enter the following commands to compile an executable from the source:

    For the current platform:

        go build github.com/asukakenji/drawing-challenge

    Cross compile for other platforms (execute one of the following commands):

        GOOS=windows GOARCH=386 go build github.com/asukakenji/drawing-challenge
        GOOS=windows GOARCH=amd64 go build github.com/asukakenji/drawing-challenge
        GOOS=darwin GOARCH=amd64 go build github.com/asukakenji/drawing-challenge
        GOOS=linux GOARCH=386 go build github.com/asukakenji/drawing-challenge
        GOOS=linux GOARCH=amd64 go build github.com/asukakenji/drawing-challenge

## Technical Diagrams

### Architecture Diagram

// Diagram //

### Package Diagram

There are 6 library packages and 1 main package, as shown in the diagram:

// Diagram //

Package common defines types and variables
which are needed by other packages in the project.

Package color defines the Color interface,
the ByteColor type which implements it, the Parser interface,
and the ByteColorParser type which implements it.

Package canvas defines the Canvas interface,
the BufferBasedCanvas interface,
and the ByteColorBuffer type which implements it.

Package command defines the Command interface,
several types which implement it, the Parser interface,
and the BasicParser type which implements it.

Package renderer defines the Renderer interface,
and the WriterRenderer type which implements it.

Package interpreter defines the Interpreter interface,
and the BasicInterpreter type which implements it.

### Class Diagram

// Diagram //

## Design Documentation

### Empty Command Behavior

If the user presses enter without entering any command, the prompt will be
printed again.

This behavior is influenced by most existing REPL (Read-Eval-Print Loop).

### New Canvas Behavior

The new canvas function creates a new canvas. If a canvas already exists, it is
destroyed and replaced by the new one.

Another option is to tell the user that a canvas is already created, and refuse
to create a new one. However, this seems not robost enough since the user needs
to quit and execute the program again to create another canvas.

### Bucket Fill Behavior

The bucket fill function fills the area enclosing (x, y). The pixels connecting
to (x, y) having the same color that at (x, y) are replaced by c.

This behavior is influenced by most existing drawing software.

## API Documentation

// TODO: Write this!
