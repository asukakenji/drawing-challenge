package bytecolor

import (
	"testing"

	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

func TestParser_ParseColor(t *testing.T) {
	parser := &Parser{Color(' ')}

	// Positive Cases
	casesPos := []struct {
		s     string
		color color.Color
	}{
		{"A", Color('A')},
		{"", Color(' ')},
	}
	for _, c := range casesPos {
		color, err := parser.ParseColor(c.s)
		if err != nil {
			t.Errorf("Case: %s, Expected: err == nil, Got: %#v", c.s, err)
		}
		if color != c.color {
			t.Errorf("Case: %s, Expected: %#v, Got: %#v", c.s, c.color, color)
		}
	}

	// Negative Cases
	casesNeg := []struct {
		s   string
		err error
	}{
		{"AA", common.ErrInvalidColor},
	}
	for _, c := range casesNeg {
		_, err := parser.ParseColor(c.s)
		if err != c.err {
			t.Errorf("Case: %s, Expected: err == nil, Got: %#v", c.s, err)
		}
	}
}
