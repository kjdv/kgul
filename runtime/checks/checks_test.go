package checks

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kjdv/kgul/testing/expects"
)

func expectFail(expect *expects.Expect, message string, fn func()) {
	defer func() {
		err := recover()

		expect.NotNil(err)
		expect.Equals(message, fmt.Sprint(err))
	}()

	fn()

	expect.Fail("unexpected success")
}

func TestCheck(t *testing.T) {
	expect := expects.New(t)

	expect.DontPanic(func() { Check(true, "") })
	expectFail(&expect, "abc3", func() { Check(false, "abc", 3) })
}

func TestCheckf1(t *testing.T) {
	expect := expects.New(t)

	expectFail(&expect, "something", func() {
		Checkf(false, "something")
	})

	expectFail(&expect, "with argument 3", func() {
		Checkf(false, "with argument %d", 3)
	})
}

func TestFail(t *testing.T) {
	expect := expects.New(t)

	expectFail(&expect, "abc3", func() { Fail("abc", 3) })
}

func TestFailf(t *testing.T) {
	expect := expects.New(t)

	expectFail(&expect, "abc", func() { Failf("abc") })
	expectFail(&expect, "with argument 3", func() { Failf("with argument %d", 3) })
}

func TestNotErr(t *testing.T) {
	expect := expects.New(t)

	expect.DontPanic(func() { NotErr(nil) })

	expectFail(&expect, "bad stuff", func() { NotErr(errors.New("bad stuff")) })

}

func TestNotErrf(t *testing.T) {
	expect := expects.New(t)

	expect.DontPanic(func() { NotErrf(nil, "") })
	expectFail(&expect, "bad stuff: with argument 3", func() { NotErrf(errors.New("bad stuff"), "%s: with argument %d", 3) })
}
