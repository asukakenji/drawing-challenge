# Drawing Challenge

The latest version of this project is available at:

https://github.com/asukakenji/drawing-challenge

## Manuals

- [User Manual](#user-manual)
- [Project Architecture Documentation](#project-architecture-documentation)
- [Technical Design Documentation](#technical-design-documentation)

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

## Project Architecture Documentation

## Technical Design Documentation
