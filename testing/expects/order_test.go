package expects

import (
	"kdv/testing/metatest"
	"testing"
)

type orderedCase struct {
	left   interface{}
	right  interface{}
	expect bool
}

func (o orderedCase) msg(idx int) []interface{} {
	m := make([]interface{}, 0)
	return append(m, "case: ", idx, ", left: ", o.left, ", right: ", o.right, ", expect: ", o.expect)
}

type myType int

func (m myType) LessThan(other interface{}) bool {
	o, ok := other.(myType)

	return ok && int(m) < int(o)
}

func TestExpect_Less(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcase := []orderedCase{
		{myType(1), myType(2), true},
		{myType(2), myType(1), false},
		{myType(1), myType(1), false},

		{1, 2, true},
		{2, 1, false},
		{1, 1, false},

		{int8(1), int8(2), true},
		{int8(2), int8(1), false},
		{int8(1), int8(1), false},

		{int16(1), int16(2), true},
		{int32(1), int32(2), true},
		{int64(1), int64(2), true},

		{uint(1), uint(2), true},
		{uint8(1), uint8(2), true},
		{uint16(1), uint16(2), true},
		{uint32(1), uint32(2), true},
		{uint64(1), uint64(2), true},

		{byte(1), byte(2), true},

		{1.0, 2.0, true},
		{2.0, 1.0, false},
		{1.0, 1.0, false},

		{float32(1), float32(2), true},
		{float64(1), float64(2), true},

		{"aaa", "bbb", true},
		{"bbb", "aaa", false},
		{"aaa", "aaa", false},
	}

	for idx, testcase := range testcase {
		e.Less(testcase.left, testcase.right, testcase.msg(idx)...)

		if testcase.expect {
			shouldSucceed(mt, t)
		} else {
			shouldFail(mt, t)
		}
	}
}

func TestExpect_LessEqual(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcase := []orderedCase{
		{myType(1), myType(2), true},
		{myType(2), myType(1), false},
		{myType(1), myType(1), true},

		{1, 2, true},
		{2, 1, false},
		{1, 1, true},

		{int8(1), int8(2), true},
		{int8(2), int8(1), false},
		{int8(1), int8(1), true},

		{int16(1), int16(2), true},
		{int32(1), int32(2), true},
		{int64(1), int64(2), true},

		{uint(1), uint(2), true},
		{uint8(1), uint8(2), true},
		{uint16(1), uint16(2), true},
		{uint32(1), uint32(2), true},
		{uint64(1), uint64(2), true},

		{byte(1), byte(2), true},

		{1.0, 2.0, true},
		{2.0, 1.0, false},
		{1.0, 1.0, true},

		{float32(1), float32(2), true},
		{float64(1), float64(2), true},

		{"aaa", "bbb", true},
		{"bbb", "aaa", false},
		{"aaa", "aaa", true},
	}

	for idx, testcase := range testcase {
		e.LessEqual(testcase.left, testcase.right, testcase.msg(idx)...)

		if testcase.expect {
			shouldSucceed(mt, t)
		} else {
			shouldFail(mt, t)
		}
	}
}

func TestExpect_GreaterEqual(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcase := []orderedCase{
		{myType(1), myType(2), false},
		{myType(2), myType(1), true},
		{myType(1), myType(1), true},

		{1, 2, false},
		{2, 1, true},
		{1, 1, true},

		{int8(1), int8(2), false},
		{int8(2), int8(1), true},
		{int8(1), int8(1), true},

		{int16(1), int16(2), false},
		{int32(1), int32(2), false},
		{int64(1), int64(2), false},

		{uint(1), uint(2), false},
		{uint8(1), uint8(2), false},
		{uint16(1), uint16(2), false},
		{uint32(1), uint32(2), false},
		{uint64(1), uint64(2), false},

		{byte(1), byte(2), false},

		{1.0, 2.0, false},
		{2.0, 1.0, true},
		{1.0, 1.0, true},

		{float32(1), float32(2), false},
		{float64(1), float64(2), false},

		{"aaa", "bbb", false},
		{"bbb", "aaa", true},
		{"aaa", "aaa", true},
	}

	for idx, testcase := range testcase {
		e.GreaterEqual(testcase.left, testcase.right, testcase.msg(idx)...)

		if testcase.expect {
			shouldSucceed(mt, t)
		} else {
			shouldFail(mt, t)
		}
	}
}

func TestExpect_Greater(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcase := []orderedCase{
		{myType(1), myType(2), false},
		{myType(2), myType(1), true},
		{myType(1), myType(1), false},

		{1, 2, false},
		{2, 1, true},
		{1, 1, false},

		{int8(1), int8(2), false},
		{int8(2), int8(1), true},
		{int8(1), int8(1), false},

		{int16(1), int16(2), false},
		{int32(1), int32(2), false},
		{int64(1), int64(2), false},

		{uint(1), uint(2), false},
		{uint8(1), uint8(2), false},
		{uint16(1), uint16(2), false},
		{uint32(1), uint32(2), false},
		{uint64(1), uint64(2), false},

		{byte(1), byte(2), false},

		{1.0, 2.0, false},
		{2.0, 1.0, true},
		{1.0, 1.0, false},

		{float32(1), float32(2), false},
		{float64(1), float64(2), false},

		{"aaa", "bbb", false},
		{"bbb", "aaa", true},
		{"aaa", "aaa", false},
	}

	for idx, testcase := range testcase {
		e.Greater(testcase.left, testcase.right, testcase.msg(idx)...)

		if testcase.expect {
			shouldSucceed(mt, t)
		} else {
			shouldFail(mt, t)
		}
	}
}

func TestExpect_NotComparable(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	type notComparable int

	e.Panics(func() { e.Less(notComparable(1), notComparable(2)) })
	shouldSucceed(mt, t)
}
