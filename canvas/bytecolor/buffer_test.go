package bytecolor

import (
	"reflect"
	"testing"

	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/color/bytecolor"
	"github.com/asukakenji/drawing-challenge/common"
)

func TestNewBuffer(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		pixels  []bytecolor.Color
	}{
		{1, 1, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' '}},
		{1, 2, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' '}},
		{1, 3, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' '}},
		{1, 4, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' '}},
		{1, 5, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' '}},
		{1, 6, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 7, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 8, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 9, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 10, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 11, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 12, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 13, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 14, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 15, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{1, 16, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' '}},
		{3, 2, bytecolor.Color(' '), bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', ' '}},
	}
	for _, c := range casesPos {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
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
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		err     error
	}{
		{0, 1, bytecolor.Color(' '), bytecolor.Color('x'), common.ErrWidthOrHeightNotPositive},
		{1, 0, bytecolor.Color(' '), bytecolor.Color('x'), common.ErrWidthOrHeightNotPositive},
	}
	for _, c := range casesNeg {
		_, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.err, err)
		}
	}
}

func TestBuffer_Dimensions(t *testing.T) {
	cases := []struct {
		w       int
		h       int
		bgColor bytecolor.Color
		fgColor bytecolor.Color
	}{
		{24, 42, bytecolor.Color(' '), bytecolor.Color('x')},
	}
	for _, c := range cases {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, err)
		}
		w, h := cnv.Dimensions()
		if w != c.w || h != c.h {
			t.Errorf("Case: (%d, %d, %#v, %#v), Expected: (%d, %d), Got: (%d, %d)", c.w, c.h, c.bgColor, c.fgColor, c.w, c.h, w, h)
		}
	}
}

// This type is created for testing purpose only
type dummyColor byte

func (c1 dummyColor) Equals(c2 color.Color) bool {
	return true
}

func TestBuffer_Set(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		x       int
		y       int
		c       color.Color
		pixels  []bytecolor.Color
	}{
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, bytecolor.Color('x'), []bytecolor.Color{'x', ' ', ' ', ' ', ' ', ' '}},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 0, bytecolor.Color('x'), []bytecolor.Color{' ', 'x', ' ', ' ', ' ', ' '}},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 1, bytecolor.Color('x'), []bytecolor.Color{' ', ' ', 'x', ' ', ' ', ' '}},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 1, bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', 'x', ' ', ' '}},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 2, bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', 'x', ' '}},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 2, bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', 'x'}},

		{3, 2, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, bytecolor.Color('x'), []bytecolor.Color{'x', ' ', ' ', ' ', ' ', ' '}},
		{3, 2, bytecolor.Color(' '), bytecolor.Color('x'), 1, 0, bytecolor.Color('x'), []bytecolor.Color{' ', 'x', ' ', ' ', ' ', ' '}},
		{3, 2, bytecolor.Color(' '), bytecolor.Color('x'), 2, 0, bytecolor.Color('x'), []bytecolor.Color{' ', ' ', 'x', ' ', ' ', ' '}},
		{3, 2, bytecolor.Color(' '), bytecolor.Color('x'), 0, 1, bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', 'x', ' ', ' '}},
		{3, 2, bytecolor.Color(' '), bytecolor.Color('x'), 1, 1, bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', 'x', ' '}},
		{3, 2, bytecolor.Color(' '), bytecolor.Color('x'), 2, 1, bytecolor.Color('x'), []bytecolor.Color{' ', ' ', ' ', ' ', ' ', 'x'}},
	}
	for _, c := range casesPos {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, err)
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
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		x       int
		y       int
		c       color.Color
		err     error
	}{
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), -1, 0, bytecolor.Color('x'), common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, -1, bytecolor.Color('x'), common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 2, 0, bytecolor.Color('x'), common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 3, bytecolor.Color('x'), common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 1, dummyColor('x'), common.ErrColorTypeNotSupported},
	}
	for _, c := range casesNeg {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, err)
		}
		err = cnv.Set(c.x, c.y, c.c)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.err, err)
		}
	}
}

func TestBuffer_At(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		x       int
		y       int
		c       bytecolor.Color
		xAt     int
		yAt     int
		cAt     bytecolor.Color
	}{
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 1, bytecolor.Color('x'), 0, 0, bytecolor.Color(' ')},
	}
	for _, c := range casesPos {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
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
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		x       int
		y       int
		c       bytecolor.Color
		xAt     int
		yAt     int
		err     error
	}{
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 1, bytecolor.Color('x'), -1, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 1, bytecolor.Color('x'), 0, -1, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 1, bytecolor.Color('x'), 2, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 1, 1, bytecolor.Color('x'), 0, 3, common.ErrPointOutsideCanvas},
	}
	for _, c := range casesNeg {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %#v, %d, %d), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x, c.y, c.c, c.xAt, c.yAt, err)
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

func TestBuffer_DrawLine(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		x1      int
		y1      int
		x2      int
		y2      int
		pixels  []bytecolor.Color
	}{
		{20, 4, bytecolor.Color(' '), bytecolor.Color('x'), 0, 1, 5, 1, []bytecolor.Color{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			'x', 'x', 'x', 'x', 'x', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 2
		{20, 4, bytecolor.Color(' '), bytecolor.Color('x'), 5, 2, 5, 3, []bytecolor.Color{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 3
		{20, 4, bytecolor.Color(' '), bytecolor.Color('x'), 5, 1, 0, 1, []bytecolor.Color{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			'x', 'x', 'x', 'x', 'x', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 2 (points in reversed order)
		{20, 4, bytecolor.Color(' '), bytecolor.Color('x'), 5, 3, 5, 2, []bytecolor.Color{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 3 (points in reverse order)
	}
	for _, c := range casesPos {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
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
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		x1      int
		y1      int
		x2      int
		y2      int
		err     error
	}{
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), -1, 0, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, -1, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 2, 0, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 3, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, -1, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, 0, -1, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, 2, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, 0, 3, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, 1, 1, common.ErrLineNotHorizontalOrVertical},
	}
	for _, c := range casesNeg {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
		}
		err = cnv.DrawLine(c.x1, c.y1, c.x2, c.y2)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, c.err, err)
		}
	}
}

func TestBuffer_DrawRect(t *testing.T) {
	// Positive Cases
	casesPos := []struct {
		w       int
		h       int
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		x1      int
		y1      int
		x2      int
		y2      int
		pixels  []bytecolor.Color
	}{
		{20, 4, bytecolor.Color(' '), bytecolor.Color('x'), 13, 0, 17, 2, []bytecolor.Color{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 4
		{20, 4, bytecolor.Color(' '), bytecolor.Color('x'), 17, 2, 13, 0, []bytecolor.Color{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}}, // Example 4 (points in reverse order)
	}
	for _, c := range casesPos {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
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
		bgColor bytecolor.Color
		fgColor bytecolor.Color
		x1      int
		y1      int
		x2      int
		y2      int
		err     error
	}{
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), -1, 0, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, -1, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 2, 0, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 3, 0, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, -1, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, 0, -1, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, 2, 0, common.ErrPointOutsideCanvas},
		{2, 3, bytecolor.Color(' '), bytecolor.Color('x'), 0, 0, 0, 3, common.ErrPointOutsideCanvas},
	}
	for _, c := range casesNeg {
		cnv, err := NewBuffer(c.w, c.h, c.bgColor, c.fgColor)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), NewBuffer returned err != nil: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, err)
		}
		err = cnv.DrawRect(c.x1, c.y1, c.x2, c.y2)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v, %#v, %d, %d, %d, %d), Expected: %#v, Got: %#v", c.w, c.h, c.bgColor, c.fgColor, c.x1, c.y1, c.x2, c.y2, c.err, err)
		}
	}
}

func TestBuffer_BucketFill(t *testing.T) {
	cnv, err := NewBuffer(20, 4, bytecolor.Color(' '), bytecolor.Color('x'))
	if err != nil {
		t.Errorf("NewBuffer returned err != nil: %#v", err)
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

	// Positive Cases
	casesPos := []struct {
		x      int
		y      int
		c      color.Color
		pixels []bytecolor.Color
	}{
		{9, 2, bytecolor.Color('o'), []bytecolor.Color{
			'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
			'x', 'x', 'x', 'x', 'x', 'x', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', ' ', ' ', ' ', 'x', 'o', 'o',
			' ', ' ', ' ', ' ', ' ', 'x', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
			' ', ' ', ' ', ' ', ' ', 'x', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o',
		}}, // Example 5
		{0, 1, bytecolor.Color('v'), []bytecolor.Color{
			'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
			'v', 'v', 'v', 'v', 'v', 'v', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', ' ', ' ', ' ', 'x', 'o', 'o',
			' ', ' ', ' ', ' ', ' ', 'v', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
			' ', ' ', ' ', ' ', ' ', 'v', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o',
		}},
		{3, 3, bytecolor.Color('o'), []bytecolor.Color{
			'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
			'v', 'v', 'v', 'v', 'v', 'v', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', ' ', ' ', ' ', 'x', 'o', 'o',
			'o', 'o', 'o', 'o', 'o', 'v', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
			'o', 'o', 'o', 'o', 'o', 'v', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o',
		}},
		{5, 1, bytecolor.Color('o'), []bytecolor.Color{
			'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
			'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', ' ', ' ', ' ', 'x', 'o', 'o',
			'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'x', 'x', 'x', 'x', 'x', 'o', 'o',
			'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o',
		}},
		{2, 2, bytecolor.Color(' '), []bytecolor.Color{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', ' ', ' ', ' ', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'x', 'x', 'x', 'x', 'x', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}},
	}
	for _, c := range casesPos {
		err = cnv.BucketFill(c.x, c.y, c.c)
		if err != nil {
			t.Errorf("Case: (%d, %d, %#v), Expected: err == nil, Got: %#v", c.x, c.y, c.c, err)
		}
		if !reflect.DeepEqual(cnv.Pixels(), c.pixels) {
			t.Errorf("Case: (%d, %d, %#v), Expected: %#v, Got: %#v", c.x, c.y, c.c, c.pixels, cnv.Pixels())
		}
	}

	// Negative Cases
	casesNeg := []struct {
		x   int
		y   int
		c   color.Color
		err error
	}{
		{-1, 0, bytecolor.Color('o'), common.ErrPointOutsideCanvas},
		{0, -1, bytecolor.Color('o'), common.ErrPointOutsideCanvas},
		{20, 0, bytecolor.Color('o'), common.ErrPointOutsideCanvas},
		{0, 4, bytecolor.Color('o'), common.ErrPointOutsideCanvas},
		{0, 0, dummyColor('o'), common.ErrColorTypeNotSupported},
	}
	for _, c := range casesNeg {
		err = cnv.BucketFill(c.x, c.y, c.c)
		if err != c.err {
			t.Errorf("Case: (%d, %d, %#v), Expected: %#v, Got: %#v", c.x, c.y, c.c, c.err, err)
		}
	}
}
