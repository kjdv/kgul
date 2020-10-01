package assert

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/kjdv/kgul/testing/expects"
)

func TestAssert(t *testing.T) {
	expect := expects.New(t)

	message := "never called"
	OnAssert(func(f Failure) {
		message = f.message
	})

	Assert(true)
	expect.Equals("never called", message)

	Assert(false) // vanilla assert
	expect.Equals("Assertion failed!", message)

	Assert(false, "with ", 3, " items")
	expect.Equals("with 3 items", message)
}

func TestAssertStackPrinter(t *testing.T) {
	expect := expects.New(t)

	message := "never called"
	OnAssert(func(f Failure) {
		message = fmt.Sprint(f)
	})

	Assert(false)
	expect.Equals(13, len(strings.Split(message, "\n")))
}

func TestAssertf(t *testing.T) {
	expect := expects.New(t)

	message := "never called"
	OnAssert(func(f Failure) {
		message = f.message
	})

	Assertf(false, "some message")
	expect.Equals("some message", message)

	Assertf(false, "with argument %d", 3)
	expect.Equals("with argument 3", message)
}

func TestFail(t *testing.T) {
	expect := expects.New(t)

	message := "never called"
	OnAssert(func(f Failure) {
		message = f.message
	})

	Fail()
	expect.Equals("Assertion failed!", message)

	Fail("with ", 3, " items")
	expect.Equals("with 3 items", message)
}

func TestFailf(t *testing.T) {
	expect := expects.New(t)

	message := "never called"
	OnAssert(func(f Failure) {
		message = f.message
	})

	Failf("some message")
	expect.Equals("some message", message)

	Failf("with argument %d", 3)
	expect.Equals("with argument 3", message)
}

func TestNotErr(t *testing.T) {
	expect := expects.New(t)

	message := "never called"
	OnAssert(func(f Failure) {
		message = f.message
	})

	NotErr(nil)
	expect.Equals("never called", message)

	NotErr(errors.New("bad stuff"))
	expect.Equals("bad stuff", message)

	NotErr(errors.New("bad stuff"), ": custom message")
	expect.Equals("bad stuff: custom message", message)

}

func TestNotErrf(t *testing.T) {
	expect := expects.New(t)

	message := "never called"
	OnAssert(func(f Failure) {
		message = f.message
	})

	NotErrf(nil, "msg")
	expect.Equals("never called", message)

	NotErrf(errors.New("bad stuff"), "some message: %s")
	expect.Equals("some message: bad stuff", message)

	NotErrf(errors.New("bad stuff"), "%s: with argument %d", 3)
	expect.Equals("bad stuff: with argument 3", message)
}

func TestPanicHandler(t *testing.T) {
	expect := expects.New(t)

	OnAssert(PanicHandler)

	expect.Panics(func() { Fail() })
}
