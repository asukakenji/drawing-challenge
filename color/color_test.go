package color

import (
	"testing"

	"github.com/asukakenji/drawing-challenge/common"
)

func TestParseByteColor(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		s     string
		color Color
	}{
		{"A", ByteColor('A')},
	}
	for _, c := range casesPos {
		color, err := ParseByteColor(c.s)
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
		{"", common.ErrEmptyColor},
		{"AA", common.ErrInvalidColor},
	}
	for _, c := range casesNeg {
		_, err := ParseByteColor(c.s)
		if err != c.err {
			t.Errorf("Case: %s, Expected: err == nil, Got: %#v", c.s, err)
		}
	}
}

// This type is created for testing purpose only
type byteColor byte

func (c1 byteColor) Equals(c2 Color) bool {
	return true
}

func TestByteColor_Equals(t *testing.T) {
	cases := []struct {
		c1     ByteColor
		c2     Color
		result bool
	}{
		{ByteColor('A'), ByteColor('A'), true},
		{ByteColor('A'), ByteColor('B'), false},
		{ByteColor(' '), ByteColor(32), true},
		{ByteColor('A'), byteColor('A'), false},
	}
	for _, c := range cases {
		result := c.c1.Equals(c.c2)
		if result != c.result {
			t.Errorf("Case: (%#v, %#v), Expected: %t, Got: %t", c.c1, c.c2, c.result, result)
		}
	}
}
