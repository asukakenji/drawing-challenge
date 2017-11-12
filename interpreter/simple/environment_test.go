package simple

import (
	"testing"

	"github.com/asukakenji/drawing-challenge/common"
)

func TestNewEnvironment(t *testing.T) {
	_, err := NewEnvironment(newCanvasFunc, &mockRenderer{})
	if err != nil {
		t.Errorf("Case #%d: Expected: err == %#v, Got: %#v", 0, nil, err)
	}

	_, err = NewEnvironment(nil, &mockRenderer{})
	if err != common.ErrNilPointer {
		t.Errorf("Case #%d: Expected: err == %#v, Got: %#v", 1, common.ErrNilPointer, err)
	}

	_, err = NewEnvironment(newCanvasFunc, nil)
	if err != common.ErrNilPointer {
		t.Errorf("Case #%d: Expected: err == %#v, Got: %#v", 2, common.ErrNilPointer, err)
	}
}

func TestEnvironment_Canvas(t *testing.T) {
	env, err := NewEnvironment(newCanvasFunc, &mockRenderer{})
	if err != nil {
		panic(err)
	}

	cnv := env.Canvas()
	if cnv != nil {
		t.Errorf("Case #%d: Expected: cnv == %#v, Got: %#v", 0, nil, cnv)
	}
}

func TestEnvironment_NewCanvas(t *testing.T) {
	env, err := NewEnvironment(newCanvasFunc, &mockRenderer{})
	if err != nil {
		panic(err)
	}

	err = env.NewCanvas(0, 0)
	if err != common.ErrWidthOrHeightNotPositive {
		t.Errorf("Case #%d: Expected: err == %#v, Got: %#v", 0, common.ErrWidthOrHeightNotPositive, err)
	}

	err = env.NewCanvas(1, 1)
	if err != nil {
		t.Errorf("Case #%d: Expected: err == %#v, Got: %#v", 1, nil, err)
	}
}

func TestEnvironment_Render(t *testing.T) {
	env, err := NewEnvironment(newCanvasFunc, &mockRenderer{})
	if err != nil {
		panic(err)
	}
	err = env.NewCanvas(1, 1)
	if err != nil {
		panic(err)
	}

	err = env.Render(env.Canvas())
	if err != nil {
		t.Errorf("Case #%d: Expected: err == %#v, Got: %#v", 0, nil, err)
	}
}

func TestEnvironment_ShouldQuit(t *testing.T) {
	env, err := NewEnvironment(newCanvasFunc, &mockRenderer{})
	if err != nil {
		panic(err)
	}

	expected := false
	got := env.ShouldQuit()
	if got != expected {
		t.Errorf("Case #%d: Expected: %t, Got: %t", 0, expected, got)
	}
}

func TestEnvironment_SetQuit(t *testing.T) {
	env, err := NewEnvironment(newCanvasFunc, &mockRenderer{})
	if err != nil {
		panic(err)
	}

	env.SetQuit()

	expected := true
	got := env.ShouldQuit()
	if got != expected {
		t.Errorf("Case #%d: Expected: %t, Got: %t", 0, expected, got)
	}
}
