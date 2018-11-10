package assert

import (
	"fmt"
	"strings"
	"testing"

	"github.com/klaasjacobdevries/kgul/testing/expects"
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

func TestPanicHandler(t *testing.T) {
	expect := expects.New(t)

	OnAssert(PanicHandler)

	expect.Panics(func() { Fail() })
}
