package writer

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	bc "github.com/asukakenji/drawing-challenge/canvas/bytecolor"
	"github.com/asukakenji/drawing-challenge/color"
	"github.com/asukakenji/drawing-challenge/color/bytecolor"
	"github.com/asukakenji/drawing-challenge/common"
)

func TestNewRenderer(t *testing.T) {
	_, err := NewRenderer(new(bytes.Buffer))
	if err != nil {
		t.Errorf("Expected: err == nil, Got: %#v", err)
	}

	_, err = NewRenderer(nil)
	if err != common.ErrNilPointer {
		t.Errorf("Expected: err == %#v, Got: %#v", common.ErrNilPointer, err)
	}
}

func TestRenderer_Render_0(t *testing.T) {
	cnv, err := bc.NewBuffer(1, 1, bytecolor.Color(' '), bytecolor.Color('x'))
	if err != nil {
		panic(err)
	}

	writer := new(bytes.Buffer)
	renderer, err := NewRenderer(writer)
	if err != nil {
		panic(err)
	}

	err = renderer.Render(cnv)
	if err != nil {
		t.Errorf("Expected: err == nil, Got: %#v", err)
	}

	expectedBytes := ([]byte)("---\n| |\n---\n\n")
	if !reflect.DeepEqual(writer.Bytes(), expectedBytes) {
		t.Errorf("Expected: %#v, Got: %#v", expectedBytes, writer.Bytes())
	}
}

func TestRenderer_Render_1(t *testing.T) {
	cnv, err := bc.NewBuffer(20, 4, bytecolor.Color(' '), bytecolor.Color('x'))
	if err != nil {
		panic(err)
	}
	err = cnv.DrawLine(0, 1, 5, 1)
	if err != nil {
		panic(err)
	}
	err = cnv.DrawLine(5, 2, 5, 3)
	if err != nil {
		panic(err)
	}
	err = cnv.DrawRect(13, 0, 17, 2)
	if err != nil {
		panic(err)
	}
	err = cnv.BucketFill(9, 2, bytecolor.Color('o'))
	if err != nil {
		panic(err)
	}

	writer := new(bytes.Buffer)
	renderer, err := NewRenderer(writer)
	if err != nil {
		panic(err)
	}

	err = renderer.Render(cnv)
	if err != nil {
		t.Errorf("Expected: err == nil, Got: %#v", err)
	}

	expectedBytes := []byte{}
	expectedBytes = append(expectedBytes, ([]byte)("----------------------\n")...)
	expectedBytes = append(expectedBytes, ([]byte)("|oooooooooooooxxxxxoo|\n")...)
	expectedBytes = append(expectedBytes, ([]byte)("|xxxxxxooooooox   xoo|\n")...)
	expectedBytes = append(expectedBytes, ([]byte)("|     xoooooooxxxxxoo|\n")...)
	expectedBytes = append(expectedBytes, ([]byte)("|     xoooooooooooooo|\n")...)
	expectedBytes = append(expectedBytes, ([]byte)("----------------------\n\n")...)

	if !reflect.DeepEqual(writer.Bytes(), expectedBytes) {
		t.Errorf("Expected: %#v, Got: %#v", expectedBytes, writer.Bytes())
	}
}

// This type is created for testing purpose only
type dummyCanvas int

func (dc dummyCanvas) Dimensions() (int, int) {
	return 0, 0
}

func (dc dummyCanvas) DrawLine(x1, y1, x2, y2 int) error {
	return nil
}

func (dc dummyCanvas) DrawRect(x1, y1, x2, y2 int) error {
	return nil
}

func (dc dummyCanvas) BucketFill(x, y int, c color.Color) error {
	return nil
}

func TestRenderer_Render_2(t *testing.T) {
	cnv := dummyCanvas(0)

	writer := new(bytes.Buffer)
	renderer, err := NewRenderer(writer)
	if err != nil {
		panic(err)
	}

	err = renderer.Render(cnv)
	if err != common.ErrCanvasNotSupported {
		t.Errorf("Expected: err == %#v, Got: %#v", common.ErrCanvasNotSupported, err)
	}
}

// This type is created for testing purpose only
type byteColor byte

func (c1 byteColor) Equals(c2 color.Color) bool {
	return true
}

func (c1 byteColor) ToByte() byte {
	return byte(c1)
}

// This type is created for testing purpose only

type byteColor2 byte

func (c1 byteColor2) Equals(c2 color.Color) bool {
	return true
}

// This type is created for testing purpose only
type anotherBufferBasedCanvas struct {
	mode   int
	width  int
	height int
	pixels []byte
}

func (abbc *anotherBufferBasedCanvas) Dimensions() (int, int) {
	return abbc.width, abbc.height
}

func (abbc *anotherBufferBasedCanvas) DrawLine(x1, y1, x2, y2 int) error {
	return nil
}

func (abbc *anotherBufferBasedCanvas) DrawRect(x1, y1, x2, y2 int) error {
	return nil
}

func (abbc *anotherBufferBasedCanvas) BucketFill(x, y int, c color.Color) error {
	return nil
}

func (abbc *anotherBufferBasedCanvas) At(x, y int) (color.Color, error) {
	b := abbc.pixels[y*abbc.width+x]
	switch abbc.mode {
	case 0:
		return bytecolor.Color(b), nil
	case 1:
		return byteColor(b), nil
	case 2:
		return byteColor2(b), nil
	default:
		return nil, errors.New("Unknown error")
	}
}

func (abbc *anotherBufferBasedCanvas) Set(x, y int, c color.Color) error {
	return nil
}

func TestRenderer_Render_3(t *testing.T) {
	cnv := &anotherBufferBasedCanvas{
		mode:   0,
		width:  3,
		height: 2,
		pixels: []byte{' ', ' ', ' ', ' ', ' ', ' '},
	}

	writer := new(bytes.Buffer)
	renderer, err := NewRenderer(writer)
	if err != nil {
		panic(err)
	}

	err = renderer.Render(cnv)
	if err != nil {
		t.Errorf("Expected: err == nil, Got: %#v", err)
	}

	cnv.mode = 1
	err = renderer.Render(cnv)
	if err != nil {
		t.Errorf("Expected: err == nil, Got: %#v", err)
	}

	cnv.mode = 2
	err = renderer.Render(cnv)
	if err != common.ErrColorNotSupported {
		t.Errorf("Expected: err == %#v, Got: %#v", common.ErrColorNotSupported, err)
	}

	cnv.mode = 3
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected: panic, Got: no panic")
		}
	}()
	renderer.Render(cnv)
}
