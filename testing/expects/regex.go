package expects

import "regexp"

func (e *Expect) Regex(pattern string, s string, msg ...interface{}) {
	r := reporter{callpoint(), msg}

	matched, err := regexp.MatchString(pattern, s)

	if err != nil {
		r.errorf(e.t, "invalid pattern: %s", err)
	} else if !matched {
		r.errorf(e.t, "expected '%s' to match pattern '%s', but didn't", s, pattern)
	}
}
