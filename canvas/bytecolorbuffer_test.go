package canvas

import (
	"reflect"
	"testing"

	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

func TestByteColorBuffer(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		pixels  []color.ByteColor
	}{
		{1, 1, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' '}},
		{1, 2, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' '}},
		{1, 3, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' '}},
		{1, 4, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' '}},
		{1, 5, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' '}},
		{1, 6, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 7, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 8, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 9, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 10, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 11, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 12, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 13, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 14, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 15, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 16, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' '}},
		{3, 2, color.ByteColor(' '), color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', ' '}},
	}
	for _, c := range casesPos {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v), Expected: err == nil, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, err)
		}
		if !reflect.DeepEqual(cnv.Pixels(), c.pixels) {
			t.Errorf("Case: (%d, %d, %#v, %#v), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.pixels, cnv.Pixels())
		}
	}

	// Negative Cases
	casesNeg := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		err     error
	}{
		{0, 1, color.ByteColor(' '), color.ByteColor('x'), common.ErrInvalidNumber},
		{1, 0, color.ByteColor(' '), color.ByteColor('x'), common.ErrInvalidNumber},
	}
	for _, c := range casesNeg {
		_, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.err, err)
		}
	}
}

func TestByteColorBuffer_Dimensions(t *testing.T) {
	cases := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
	}{
		{24, 42, color.ByteColor(' '), color.ByteColor('x')},
	}
	for _, c := range cases {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v), Newcolor.ByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, err)
		}
		w, h := cnv.Dimensions()
		if w != c.w || h != c.h {
			t.Errorf("Case: (%d, %d, %#v, %#v), Expected: (%d, %d), Got: (%d, %d)", c.w, c.h, c.bgColor, c.fgColor, c.w, c.h, w, h)
		}
	}
}

// This type is created for testing purpose only
type byteColor byte

func (c1 byteColor) Equals(c2 color.Color) bool {
	return true
}

func TestByteColorBuffer_Set(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		x       int
		y       int
		c       color.Color
		pixels  []color.ByteColor
	}{
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, color.ByteColor('x'), []color.ByteColor{'x', ' ', ' ', ' ', ' ', ' '}},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 0, color.ByteColor('x'), []color.ByteColor{' ', 'x', ' ', ' ', ' ', ' '}},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 1, color.ByteColor('x'), []color.ByteColor{' ', ' ', 'x', ' ', ' ', ' '}},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', 'x', ' ', ' '}},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 2, color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', 'x', ' '}},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 2, color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', 'x'}},

		{3, 2, color.ByteColor(' '), color.ByteColor('x'), 0, 0, color.ByteColor('x'), []color.ByteColor{'x', ' ', ' ', ' ', ' ', ' '}},
		{3, 2, color.ByteColor(' '), color.ByteColor('x'), 1, 0, color.ByteColor('x'), []color.ByteColor{' ', 'x', ' ', ' ', ' ', ' '}},
		{3, 2, color.ByteColor(' '), color.ByteColor('x'), 2, 0, color.ByteColor('x'), []color.ByteColor{' ', ' ', 'x', ' ', ' ', ' '}},
		{3, 2, color.ByteColor(' '), color.ByteColor('x'), 0, 1, color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', 'x', ' ', ' '}},
		{3, 2, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', 'x', ' '}},
		{3, 2, color.ByteColor(' '), color.ByteColor('x'), 2, 1, color.ByteColor('x'), []color.ByteColor{' ', ' ', ' ', ' ', ' ', 'x'}},
	}
	for _, c := range casesPos {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), Newcolor.ByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, err)
		}
		err = cnv.Set(c.x, c.y, c.c)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), Expected: err == nil, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, err)
		}
		if !reflect.DeepEqual(cnv.Pixels(), c.pixels) {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.pixels, cnv.Pixels())
		}
	}

	// Negative Cases
	casesNeg := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		x       int
		y       int
		c       color.Color
		err     error
	}{
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), -1, 0, color.ByteColor('x'), common.ErrInvalidNumber},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, -1, color.ByteColor('x'), common.ErrInvalidNumber},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 2, 0, color.ByteColor('x'), common.ErrInvalidNumber},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 3, color.ByteColor('x'), common.ErrInvalidNumber},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, byteColor('x'), common.ErrInvalidColor},
	}
	for _, c := range casesNeg {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), Newcolor.ByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, err)
		}
		err = cnv.Set(c.x, c.y, c.c)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.err, err)
		}
	}
}

func TestByteColorBuffer_At(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		x       int
		y       int
		c       color.ByteColor
		xAt     int
		yAt     int
		cAt     color.ByteColor
	}{
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), 0, 0, color.ByteColor(' ')},
	}
	for _, c := range casesPos {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), Newcolor.ByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
		}
		err = cnv.Set(c.x, c.y, c.c)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), cnv.Set returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
		}
		cAt, err := cnv.At(c.xAt, c.yAt)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), Expected: err == nil, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
		}
		if cAt != c.cAt {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, c.cAt, cAt)
		}
	}

	// Negative Cases
	casesNeg := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		x       int
		y       int
		c       color.ByteColor
		xAt     int
		yAt     int
		err     error
	}{
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), -1, 0, common.ErrInvalidNumber},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), 0, -1, common.ErrInvalidNumber},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), 2, 0, common.ErrInvalidNumber},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), 0, 3, common.ErrInvalidNumber},
	}
	for _, c := range casesNeg {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), Newcolor.ByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
		}
		err = cnv.Set(c.x, c.y, c.c)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), cnv.Set returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
		}
		_, err = cnv.At(c.xAt, c.yAt)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, c.err, err)
		}
	}
}
