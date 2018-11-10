package expects

import (
	"fmt"
	"kdv/testing/metatest"
	"testing"
)

type comparisonCase struct {
	expect interface{}
	actual interface{}
}

func (c comparisonCase) msg(idx int) []interface{} {
	m := make([]interface{}, 0)
	return append(m, "case: ", idx, " expect: ", c.expect, " actual: ", c.actual)
}

func TestExpect_Equals(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcases := []comparisonCase{
		{true, true},
		{1, 1},
		{uint(2), uint(2)},
		{3.3, 3.3},
		{float32(4.4), float32(4.4)},
		{float64(5.5), float64(5.5)},
		{complex(6.6, 6.6), complex(6.6, 6.6)},
		{"seven", "seven"},
		{[2]int{8, 9}, [2]int{8, 9}},
		{struct{ an int }{10}, struct{ an int }{10}},
		{nil, nil},
	}

	for idx, testcase := range testcases {
		e.Equals(testcase.expect, testcase.actual, testcase.msg(idx)...)
		shouldSucceed(mt, t)

		e.NotEquals(testcase.expect, testcase.actual, testcase.msg(idx)...)
		shouldFail(mt, t)
	}
}

func TestExpect_NotEquals(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcases := []comparisonCase{
		{true, false},
		{1, 2},
		{uint(2), uint(3)},
		{3.3, 3.4},
		{float32(4.4), float32(4.5)},
		{float64(5.5), float64(5.6)},
		{complex(6.6, 6.6), complex(6.6, 6.7)},
		{"seven", "eight"},
		{[2]int{8, 9}, [2]int{8, 10}},
		{struct{ an int }{10}, struct{ an int }{11}},

		{true, 1},
		{1, uint(1)},
		{float32(2.2), float64(2.2)},
	}

	for idx, testcase := range testcases {
		e.Equals(testcase.expect, testcase.actual, testcase.msg(idx)...)
		shouldFail(mt, t)

		e.NotEquals(testcase.expect, testcase.actual, testcase.msg(idx)...)
		shouldSucceed(mt, t)
	}
}

func TestExpect_AlmostEquals(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcases := []struct {
		expect   float64
		actual   float64
		epsilon  float64
		succeeds bool
	}{
		{1.0, 1.0, 0.0, true},
		{1.0, 1.1, 0.0, false},
		{1.0, 1.1, 0.2, true},
		{1.0, 1.1, 0.05, false},
	}

	for idx, tc := range testcases {
		e.AlmostEqual(tc.expect, tc.actual, tc.epsilon, fmt.Sprintf("case %d", idx))

		if tc.succeeds {
			shouldSucceed(mt, t)
		} else {
			shouldFail(mt, t)
		}
	}
}

func TestExpect_NotAlmostEquals(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcases := []struct {
		expect   float64
		actual   float64
		epsilon  float64
		succeeds bool
	}{
		{1.0, 1.0, 0.1, false},
		{1.0, 1.1, 0.01, true},
		{1.0, 1.1, 0.2, false},
		{1.0, 1.1, 0.05, true},
	}

	for idx, tc := range testcases {
		e.NotAlmostEqual(tc.expect, tc.actual, tc.epsilon, fmt.Sprintf("case %d", idx))

		if tc.succeeds {
			shouldSucceed(mt, t)
		} else {
			shouldFail(mt, t)
		}
	}
}
