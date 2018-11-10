package expects

func (e *Expect) True(condition bool, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if !condition {
		r.errorf(e.t, "expected true, but was false")
	}
}

func (e *Expect) False(condition bool, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if condition {
		r.errorf(e.t, "expected false, but was true")
	}
}

func (e *Expect) Fail(msg ...interface{}) {
	r := reporter{callpoint(), msg}

	r.errorf(e.t, "Fail")
}

func (e *Expect) IsNil(value interface{}, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if value != nil {
		r.errorf(e.t, "expected nil but was "+printVariable(value))
	}
}

func (e *Expect) NotNil(value interface{}, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if value == nil {
		r.errorf(e.t, "unexpectedly nil")
	}
}
