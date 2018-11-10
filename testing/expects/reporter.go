package expects

import (
	"fmt"

	"github.com/klaasjacobdevries/kgul/testing/metatest"
)

func makeBody(msg ...interface{}) string {
	if len(msg) == 0 {
		return ""
	}
	return "\n" + fmt.Sprint(msg...)
}

func printVariable(variable interface{}) string {
	return fmt.Sprintf("(%T) %+v", variable, variable)
}

type reporter struct {
	callpoint string
	msg       []interface{}
}

func (r reporter) errorf(t metatest.Tester, format string, args ...interface{}) {
	head := fmt.Sprintf(format, args...)
	body := makeBody(r.msg...)
	tail := "\nfrom " + r.callpoint

	t.Error(head + body + tail)
}

func (r reporter) comparison(t metatest.Tester, head string, expect interface{}, actual interface{}) {
	body := "\nexpected: " + printVariable(expect) + "\nactual:   " + printVariable(actual)

	body += makeBody(r.msg...)
	tail := "\nfrom " + r.callpoint

	t.Error(head + body + tail)
}

func (r reporter) ordering(t metatest.Tester, head string, left interface{}, right interface{}) {
	body := "\nleft:  " + printVariable(left) + "\nright: " + printVariable(right)

	body += makeBody(r.msg...)
	tail := "\nfrom " + r.callpoint

	t.Error(head + body + tail)
}

func (r reporter) mismatch(t metatest.Tester, head string, value interface{}, matcher interface{}) {
	body := "\nvalue:   " + printVariable(value) + "\nmatcher: " + printVariable(matcher)
	body += makeBody(r.msg...)
	tail := "\nfrom " + r.callpoint

	t.Error(head + body + tail)
}
