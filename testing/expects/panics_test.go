package expects

import (
	"kdv/testing/metatest"
	"testing"
)

func TestExpect_Panics(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	fn := func(a, b int) int {
		return a / b
	}

	e.Panics(func() { fn(1, 0) }, "division by zero panics")
	shouldSucceed(mt, t)

	e.Panics(func() { fn(1, 2) }, "this shouldn't panic")
	shouldFail(mt, t)

	// this is meta...
	var nilFunc func() = nil
	e.Panics(nilFunc, "not callable")
	shouldSucceed(mt, t)
}

func TestExpect_DontPanic(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	fn := func(a, b int) int {
		return a / b
	}

	e.DontPanic(func() { fn(1, 2) }, "does not pannic")
	shouldSucceed(mt, t)

	e.DontPanic(func() { fn(1, 0) }, "does pannic")
	shouldFail(mt, t)

	var nilFunc func() = nil
	e.DontPanic(nilFunc)
	shouldFail(mt, t)
}
