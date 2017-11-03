package color

import "testing"

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
