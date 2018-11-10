package expects

import (
	"kdv/testing/metatest"
	"testing"
)

func TestExpect_True(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	e.True(true, "this is true")
	shouldSucceed(mt, t)

	e.True(false, "this is false")
	shouldFail(mt, t)
}

func TestExpect_False(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	e.False(true, "this is true")
	shouldFail(mt, t)

	e.False(false, "this is false")
	shouldSucceed(mt, t)
}

func TestExpect_Fail(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	e.Fail("failure")
	shouldFail(mt, t)
}

func TestExpect_IsNil(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	e.IsNil(nil, "this is nil")
	shouldSucceed(mt, t)

	e.IsNil(0, "this is not nil")
	shouldFail(mt, t)
}

func TestExpect_NotNil(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	e.NotNil(nil, "this is nil")
	shouldFail(mt, t)

	e.NotNil(0, "this is not nil")
	shouldSucceed(mt, t)
}
