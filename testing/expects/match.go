package expects

type Matcher interface {
	Match(value interface{}) bool
}

func (e *Expect) That(value interface{}, matches Matcher, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	if !matches.Match(value) {
		r.mismatch(e.t, "matching failed", value, matches)
	}
}
