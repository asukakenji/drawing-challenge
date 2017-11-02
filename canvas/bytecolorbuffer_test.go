package canvas

import (
	"reflect"
	"testing"

	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/common"
)

func TestNewByteColorBuffer(t *testing.T) {
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
		{0, 1, color.ByteColor(' '), color.ByteColor('x'), common.ErrWidthOrHeightNotPositive},
		{1, 0, color.ByteColor(' '), color.ByteColor('x'), common.ErrWidthOrHeightNotPositive},
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
			t.Errorf("Case: (%d, %d, %#v, %#v), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, err)
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
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, err)
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
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), -1, 0, color.ByteColor('x'), common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, -1, color.ByteColor('x'), common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 2, 0, color.ByteColor('x'), common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 3, color.ByteColor('x'), common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, byteColor('x'), common.ErrColorTypeNotSupported},
	}
	for _, c := range casesNeg {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, err)
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
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
		}
		err = cnv.Set(c.x, c.y, c.c)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), Set returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
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
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), -1, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), 0, -1, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), 2, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 1, 1, color.ByteColor('x'), 0, 3, common.ErrPointOutsideCanvas},
	}
	for _, c := range casesNeg {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
		}
		err = cnv.Set(c.x, c.y, c.c)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), Set returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
		}
		_, err = cnv.At(c.xAt, c.yAt)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, c.err, err)
		}
	}
}

func TestByteColorBuffer_DrawLine(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		x1      int
		y1      int
		x2      int
		y2      int
		pixels  []color.ByteColor
	}{
		{20, 4, color.ByteColor(' '), color.ByteColor('x'), 0, 1, 5, 1, []color.ByteColor{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			'x', 'x', 'x', 'x', 'x', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 2
		{20, 4, color.ByteColor(' '), color.ByteColor('x'), 5, 2, 5, 3, []color.ByteColor{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 3
		{20, 4, color.ByteColor(' '), color.ByteColor('x'), 5, 1, 0, 1, []color.ByteColor{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			'x', 'x', 'x', 'x', 'x', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 2 (points in reversed order)
		{20, 4, color.ByteColor(' '), color.ByteColor('x'), 5, 3, 5, 2, []color.ByteColor{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 3 (points in reverse order)
	}
	for _, c := range casesPos {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
		}
		err = cnv.DrawLine(c.x1, c.y1, c.x2, c.y2)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), Expected: err == nil, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
		}
		if !reflect.DeepEqual(cnv.Pixels(), c.pixels) {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, c.pixels, cnv.Pixels())
		}
	}

	// Negative Cases
	casesNeg := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		x1      int
		y1      int
		x2      int
		y2      int
		err     error
	}{
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), -1, 0, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, -1, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 2, 0, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 3, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, -1, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, 0, -1, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, 2, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, 0, 3, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, 1, 1, common.ErrLineNotHorizontalOrVertical},
	}
	for _, c := range casesNeg {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
		}
		err = cnv.DrawLine(c.x1, c.y1, c.x2, c.y2)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, c.err, err)
		}
	}
}

func TestByteColorBuffer_DrawRect(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		x1      int
		y1      int
		x2      int
		y2      int
		pixels  []color.ByteColor
	}{
		{20, 4, color.ByteColor(' '), color.ByteColor('x'), 13, 0, 17, 2, []color.ByteColor{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 4
		{20, 4, color.ByteColor(' '), color.ByteColor('x'), 17, 2, 13, 0, []color.ByteColor{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 4 (points in reverse order)
	}
	for _, c := range casesPos {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
		}
		err = cnv.DrawRect(c.x1, c.y1, c.x2, c.y2)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), Expected: err == nil, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
		}
		if !reflect.DeepEqual(cnv.Pixels(), c.pixels) {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, c.pixels, cnv.Pixels())
		}
	}

	// Negative Cases
	casesNeg := []struct {
		w       int
		h       int
		bgColor color.ByteColor
		fgColor color.ByteColor
		x1      int
		y1      int
		x2      int
		y2      int
		err     error
	}{
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), -1, 0, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, -1, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 2, 0, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 3, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, -1, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, 0, -1, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, 2, 0, common.ErrPointOutsideCanvas},
		{2, 3, color.ByteColor(' '), color.ByteColor('x'), 0, 0, 0, 3, common.ErrPointOutsideCanvas},
	}
	for _, c := range casesNeg {
		cnv, err := NewByteColorBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), NewByteColorBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
		}
		err = cnv.DrawRect(c.x1, c.y1, c.x2, c.y2)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, c.err, err)
		}
	}
}

func TestByteColorBuffer_BucketFill(t *testing.T) {
	cnv, err := NewByteColorBuffer(20, 4, color.ByteColor(' '), color.ByteColor('x'))
	if err != nil {
		t.Errorf("NewByteColorBuffer returned err != nil: %#v", err)
	}
	err = cnv.DrawLine(0, 1, 5, 1) // Example 2
	if err != nil {
		t.Errorf("DrawLine returned err != nil: %#v", err)
	}
	err = cnv.DrawLine(5, 2, 5, 3) // Example 3
	if err != nil {
		t.Errorf("DrawLine returned err != nil: %#v", err)
	}
	err = cnv.DrawRect(13, 0, 17, 2) // Example 4
	if err != nil {
		t.Errorf("DrawRect returned err != nil: %#v", err)
	}

	pixels0 := []color.ByteColor{
		'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
		'x', 'x', 'x', 'x', 'x', 'x', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', ' ', ' ', ' ', 'x', 'o', 'o',
		' ', ' ', ' ', ' ', ' ', 'x', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
		' ', ' ', ' ', ' ', ' ', 'x', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o',
	}
	err = cnv.BucketFill(9, 2, color.ByteColor('o')) // Example 5
	if err != nil {
		t.Errorf("BucketFill returned err != nil: %#v", err)
	}
	if !reflect.DeepEqual(cnv.Pixels(), pixels0) {
		t.Errorf("Case 0: Expected: %#v, Got: %#v", pixels0, cnv.Pixels())
	}

	err1 := common.ErrPointOutsideCanvas
	err = cnv.BucketFill(-1, 0, color.ByteColor('o'))
	if err != err1 {
		t.Errorf("Case 1a: Expected: %#v, Got: %#v", err1, err)
	}
	err = cnv.BucketFill(0, -1, color.ByteColor('o'))
	if err != err1 {
		t.Errorf("Case 1b: Expected: %#v, Got: %#v", err1, err)
	}
	err = cnv.BucketFill(20, 0, color.ByteColor('o'))
	if err != err1 {
		t.Errorf("Case 1c: Expected: %#v, Got: %#v", err1, err)
	}
	err = cnv.BucketFill(0, 4, color.ByteColor('o'))
	if err != err1 {
		t.Errorf("Case 1d: Expected: %#v, Got: %#v", err1, err)
	}

	err2 := common.ErrColorTypeNotSupported
	err = cnv.BucketFill(0, 0, byteColor('o'))
	if err != err2 {
		t.Errorf("Case 2: Expected: %#v, Got: %#v", err2, err)
	}

	// TODO: Add more cases
	// TODO: Rewrite the algorithm
}
