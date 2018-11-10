package expects

func (e *Expect) Panics(callable func(), msg ...interface{}) {
	r := reporter{callpoint(), msg}

	defer func() {
		recover()
	}()

	callable()

	// if we reach this, panic did not happen
	r.errorf(e.t, "expected to panic, but this didn't happen")
}

func (e *Expect) DontPanic(callable func(), msg ...interface{}) {
	r := reporter{callpoint(), msg}

	defer func() {
		if rc := recover(); rc != nil {
			r.errorf(e.t, "function expected not to panic, but it did")
		}
	}()

	callable()
}
