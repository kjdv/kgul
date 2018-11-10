package expects

import (
	"fmt"
	"reflect"
)

func (e *Expect) Equals(expect interface{}, actual interface{}, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if !reflect.DeepEqual(expect, actual) {
		r.comparison(e.t, "expected != actual", expect, actual)
	}
}

func (e *Expect) NotEquals(expect interface{}, actual interface{}, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if reflect.DeepEqual(expect, actual) {
		r.comparison(e.t, "expected == actual", expect, actual)
	}
}

func abs(f float64) float64 {
	if f < 0.0 {
		return -f
	}
	return f
}

func absdiff(a float64, b float64) float64 {
	d := a - b
	return abs(d)
}

func (e *Expect) AlmostEqual(expect float64, actual float64, epsilon float64, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	ad := absdiff(expect, actual)
	if ad > epsilon {
		m := fmt.Sprintf("abs(expect - actual) > %v", epsilon)
		r.comparison(e.t, m, expect, actual)
	}
}

func (e *Expect) NotAlmostEqual(expect float64, actual float64, epsilon float64, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	ad := absdiff(expect, actual)
	if ad < epsilon {
		m := fmt.Sprintf("abs(expect - actual) < %v", epsilon)
		r.comparison(e.t, m, expect, actual)
	}
}
