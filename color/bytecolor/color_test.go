package bytecolor

import (
	"testing"

	"github.com/asukakenji/drawing-challenge/color"
)

// This type is created for testing purpose only
type dummyColor byte

func (c1 dummyColor) Equals(c2 color.Color) bool {
	return true
}

func TestColor_Equals(t *testing.T) {
	cases := []struct {
		c1     Color
		c2     color.Color
		result bool
	}{
		{Color('A'), Color('A'), true},
		{Color('A'), Color('B'), false},
		{Color(' '), Color(32), true},
		{Color('A'), dummyColor('A'), false},
	}
	for _, c := range cases {
		result := c.c1.Equals(c.c2)
		if result != c.result {
			t.Errorf("Case: (%#v, %#v), Expected: %t, Got: %t", c.c1, c.c2, c.result, result)
		}
	}
}
