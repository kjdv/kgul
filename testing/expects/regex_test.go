package expects

import (
	"kdv/testing/metatest"
	"testing"
)

func TestExpect_Regex(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	testcases := []struct {
		pattern string
		actual  string
		expect  bool
	}{
		{"same", "same", true},
		{"a(b", "invalid pattern", false},
		{"s.*", "same", true},
		{"z.*", "same", false},
	}

	for idx, testcase := range testcases {
		e.Regex(testcase.pattern, testcase.actual, "case: ", idx, ", string: ", testcase.actual, ", pattern: ", testcase.pattern)

		if testcase.expect {
			shouldSucceed(mt, t)
		} else {
			shouldFail(mt, t)
		}
	}
}
