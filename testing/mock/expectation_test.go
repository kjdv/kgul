package mock

import (
	"kdv/testing/expects"
	"testing"
)

func TestExpectation_Match(t *testing.T) {
	expect := expects.New(t)

	exp := newExpectation(
		deriveSignature(func(int, string, int) {}),
		[]Matcher{Eq(1), Eq("test"), Ne(2)})

	_, ok, _ := exp.match(asValues(1, "test", 1))
	expect.True(ok)

	_, ok, msg := exp.match(asValues(2, "test", 1))
	expect.False(ok)
	expect.NotEquals("", msg)

	_, ok, msg = exp.match(asValues(1, "no test", 1))
	expect.False(ok)
	expect.NotEquals("", msg)

	_, ok, msg = exp.match(asValues(1, "test", 2))
	expect.False(ok)
	expect.NotEquals("", msg)

	_, ok, msg = exp.match(asValues(1, "test", 1, 2, 3))
	expect.False(ok)
	expect.NotEquals("", msg)

	_, ok, msg = exp.match(asValues(1, "test"))
	expect.False(ok)
	expect.NotEquals("", msg)
}

func TestExpectation_verify(t *testing.T) {
	expect := expects.New(t)

	exp := newExpectation(
		deriveSignature(func(int) {}),
		[]Matcher{NotNil()})

	ok, msg := exp.verify()
	expect.False(ok)
	expect.NotEquals("", msg)
	expect.Equals(0, exp.NumMatches())

	exp.match(asValues(0))
	ok, _ = exp.verify()
	expect.True(ok)
	expect.Equals(1, exp.NumMatches())

	exp.match(asValues(1))
	ok, msg = exp.verify()
	expect.True(ok, "default cardinality should be AtLeast(1) (", msg, ")")
	expect.Equals(2, exp.NumMatches())
}

func TestExpectation_NTimes_Never(t *testing.T) {
	expect := expects.New(t)

	exp := newExpectation(
		deriveSignature(func(interface{}) {}), []Matcher{NewAny()}).
		NTimes(0)
	expect.True(exp.verify())

	exp.match(asValues(nil))
	expect.False(exp.verify())
}

func TestExpectation_NTimes_Exactly(t *testing.T) {
	expect := expects.New(t)

	exp := newExpectation(
		deriveSignature(func(interface{}) {}), []Matcher{NewAny()}).
		NTimes(2)

	expect.False(exp.verify())

	exp.match(asValues(nil))
	expect.False(exp.verify())

	exp.match(asValues(nil))
	expect.True(exp.verify())

	exp.match(asValues(nil))
	expect.False(exp.verify())
}

func TestExpectation_Times(t *testing.T) {
	expect := expects.New(t)

	exp := newExpectation(
		deriveSignature(func(interface{}) {}), []Matcher{NewAny()}).
		Times(NewAtMost(2))

	expect.True(exp.verify())

	exp.match(asValues(nil))
	expect.True(exp.verify())

	exp.match(asValues(nil))
	expect.True(exp.verify())

	exp.match(asValues(nil))
	expect.False(exp.verify())
}

func TestExpectation_SideEffect(t *testing.T) {
	expect := expects.New(t)

	called := 0
	se := func(a, b int) {
		called++
	}

	exp := newExpectation(
		deriveSignature(func(int, int) {}), []Matcher{NewAny(), NewAny()}).
		SideEffect(se)
	exp.match(asValues(1, 2))
	exp.match(asValues(3, 4))

	expect.Equals(2, called)
}

func TestExpectation_SideEffect_withPointer(t *testing.T) {
	expect := expects.New(t)

	called := 0
	se := func(a, b int) {
		called++
	}

	exp := newExpectation(
		deriveSignature(func(int, int) {}), []Matcher{NewAny(), NewAny()}).
		SideEffect(&se)
	exp.match(asValues(1, 2))
	exp.match(asValues(3, 4))

	expect.Equals(2, called)
}

func TestExpectation_InvalidSideEffect(t *testing.T) {
	expect := expects.New(t)

	se := func(a, b, c int) {}

	exp := newExpectation(
		deriveSignature(func(int, int) {}), []Matcher{NewAny(), NewAny()})

	expect.Panics(func() { exp.SideEffect(se) })
}

func TestExpectation_Returns(t *testing.T) {
	expect := expects.New(t)

	exp := newExpectation(
		deriveSignature(func() (int, int, int) { return 0, 0, 0 }), []Matcher{}).
		Returns(2, 3, 4)
	rv, ok, _ := exp.match(asValues())
	expect.True(ok)

	expect.Equals(3, len(rv))
	expect.Equals(2, asInterface(rv[0]))
	expect.Equals(3, asInterface(rv[1]))
	expect.Equals(4, asInterface(rv[2]))
}

func TestExpectation_ReturnsInvalid(t *testing.T) {
	expect := expects.New(t)

	exp := newExpectation(
		deriveSignature(func() (int, int, int) { return 0, 0, 0 }), []Matcher{})

	expect.Panics(func() { exp.Returns(1, 2) })
	expect.Panics(func() { exp.Returns(1, 2, 3, 4) })
	expect.Panics(func() { exp.Returns(1, 2, "3") })
}

func TestExpectation_String(t *testing.T) {
	expect := expects.New(t)

	exp := newExpectation(
		deriveSignature(func() {}), []Matcher{Eq(1), NotNil()})
	expect.Equals(
		"expectation with arguments [Eq(int=1) NotNil()]\n"+
			"to be called at least 1 time",
		doString(exp))
}
