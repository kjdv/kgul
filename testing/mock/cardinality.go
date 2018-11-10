package mock

import "fmt"

func fmtTime(n int) string {
	if n == 1 {
		return "1 time"
	}
	return fmt.Sprintf("%d times", n)
}

type Cardinality interface {
	Compare(ncalls int) (ok bool, msg string)
}

type Never struct {
}

func NewNever() Never {
	return Never{}
}

func (n Never) Compare(ncalls int) (ok bool, msg string) {
	if ncalls > 0 {
		return false, "expected to be never called, but was called " + fmtTime(ncalls)
	}
	return true, ""
}

func (n Never) String() string {
	return "never called"
}

type Exactly struct {
	times int
}

func NewExactly(n int) Exactly {
	return Exactly{n}
}

func (e Exactly) Compare(ncalls int) (ok bool, msg string) {
	if e.times != ncalls {
		return false, fmt.Sprintf("expected %s, but was called %s", fmtTime(e.times), fmtTime(ncalls))
	}
	return true, ""
}

func (e Exactly) String() string {
	return "called exactly " + fmtTime(e.times)
}

type AtLeast struct {
	times int
}

func NewAtLeast(n int) AtLeast {
	return AtLeast{n}
}

func (al AtLeast) Compare(ncalls int) (ok bool, msg string) {
	if !(ncalls >= al.times) {
		return false, fmt.Sprintf("not called enough (%d < %d)", ncalls, al.times)
	}
	return true, ""
}

func (al AtLeast) String() string {
	return "called at least " + fmtTime(al.times)
}

type AtMost struct {
	times int
}

func NewAtMost(n int) AtMost {
	return AtMost{n}
}

func (am AtMost) Compare(ncalls int) (ok bool, msg string) {
	if !(ncalls <= am.times) {
		return false, fmt.Sprintf("called too much (%d > %d)", ncalls, am.times)
	}
	return true, ""
}

func (am AtMost) String() string {
	return "called at most " + fmtTime(am.times)
}
