package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	inputText := `
AA 1 2 3 4
L 1 2 3 4
C 20 4
L 1 2 6 2
L 6 3 6 4
R 14 1 18 3
B 10 3 o
`
	input = strings.NewReader(inputText)

	// Neg0
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Case #0: Expected panic")
			} else {
				bgColorString = DefaultBGColorString
			}
		}()
		bgColorString = "AA"
		main()
	}()

	// Neg1
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Case #1: Expected panic")
			} else {
				fgColorString = DefaultFGColorString
			}
		}()
		fgColorString = "AA"
		main()
	}()

	// Neg2
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Case #2: Expected panic")
			} else {
				output = new(bytes.Buffer)
			}
		}()
		output = nil
		main()
	}()

	// Pos
	main()
}
