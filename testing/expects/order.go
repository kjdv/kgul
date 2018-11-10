package expects

type Ordered interface {
	LessThan(other interface{}) bool
}

func less(left interface{}, right interface{}) bool {
	if left, ok := left.(Ordered); ok {
		return left.LessThan(right)
	}

	switch left := left.(type) {

	case int:
		return left < right.(int)
	case int8:
		return left < right.(int8)
	case int16:
		return left < right.(int16)
	case int32:
		return left < right.(int32)
	case int64:
		return left < right.(int64)

	case uint:
		return left < right.(uint)
	case uint8:
		return left < right.(uint8)
	case uint16:
		return left < right.(uint16)
	case uint32:
		return left < right.(uint32)
	case uint64:
		return left < right.(uint64)

	case float32:
		return left < right.(float32)
	case float64:
		return left < right.(float64)

	case string:
		return left < right.(string)

	default:
		panic("can not compare " + printVariable(left) + " to " + printVariable(right))
	}
}

func (e *Expect) Less(left interface{}, right interface{}, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if !less(left, right) {
		r.ordering(e.t, "left >= right", left, right)
	}
}

func (e *Expect) LessEqual(left interface{}, right interface{}, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if less(right, left) {
		r.ordering(e.t, "left > right", left, right)
	}
}

func (e *Expect) Greater(left interface{}, right interface{}, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if !less(right, left) {
		r.ordering(e.t, "left <= right", left, right)
	}
}

func (e *Expect) GreaterEqual(left interface{}, right interface{}, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if less(left, right) {
		r.ordering(e.t, "left < right", left, right)
	}
}
