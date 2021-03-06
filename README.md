# kgul
Klaas' Go Utility library

A small set of assorted utilities to help with other Go projects.

## Assert

Makes it possible to use assertions, for example:

```go
import "github.com/kjdv/kgul/runtime/assert"

func divide(a, b int) int {
	assert.Assert(b != 0, "Can not divide by zero")
	return a / b
}
```

This also comes in flavours, there is an `Assertf()` for formatted assertion messages in `Fail()` and `Failf()` to unconditionally fail, useful when flow-control rather than an explicit condition signals failure.

Default behaviour is to print the message and a stack trace to stderr and exit the process with a non-zero exit code. By calling `OnAssert()` a different failure handler can be installed:

```go
import "github.com/kjdv/kgul/runtime/assert"

func setPanicHandler() {
  assert.OnAssert(assert.PanicHandler) // the provided PanicHandler panics instead of exiting
}

func setNoopHandler() {
  // if you want to live dangerously you can, set a handler that does nothing
  noop := func(assert.Failure) {}
  assert.OnAssert(noop)
}
```

## Logging

A slightly more configurable logger than the standard one. You can log by creating loggers:

```go
import "github.com/kjdv/kgul/runtime/logging"

func doLogging() {
  l := logging.New("logger.C")
  l.SetLevel(logging.Debug) // enable logging at debug level
  l.Debug("Different")
  l.Info("Loglevels")
  l.Warning("Can be")
  l.Error("Used")
}
```

This will print the log to stdout by default, or to a file passed by the `-logFile` flag. The loglevel can be set either at the level of the individual loggers using `logger.SetLevel()` or globally by passing the `-logLevel` flag.

The actual output happens on a separate goroutine so i/o operations do not block the main program.

## Test expectations

The built-in unit testing functionality for Go is nice, but very very spartan. The expects library should enable you to write tests without reverting to explicit flow control.

```go

import "github.com/kjdv/kgul/testing/expects"

func TestMyExpect(t *testing.T) {
	expect := expects.New(t)

	expect.True(2+2 == 5) // prints a friendly error message plus the file and line number of the failure
	expect.Equals(5, 2+2) // prints
	//  expected != actual
	//  expected: (int) 5
	//  actual:   (int) 4

	expect.Equals(struct{ a string }{"foo"}, struct{ a string }{"bar"} // complex structures supported, this fails with:
	//  expected != actual
	//  expected: (struct { a string }) {a:foo}
	//  actual:   (struct { a string }) {a:bar}
}
```

Other  expectations supported include `Fail()` (non-conditional, based on flow control), `IsNil()` and `IsNotNil()`, `AlmostEqual()` and `AlmostNotEqual()` (for floating-point comparisons),  `Less()`, `LessEqual()`, `Greater()` and `GreaterEqual()`, `Panics()`, `Regex()` for string matching and a generic `That()` to match for a custom matcher (c.f. Mock)

## Mock

Mocking library to accompany expects. This is in the style of C++ gmock, though admittingly in the more dynamic language that Go is it is a less attractive tradeof between usefulness and complexity. Often a simple hand-written stub is the better option.

This allows you to write tests like this:

```go
import (
	"github.com/kjdv/kgul/testing/expects"
	"github.com/kjdv/kgul/testing/metatest"
)

func TestMyMock(t *testing.T) {
	expect := expects.New(t)
	mock := New(t)

	// track a side effect
	called := 0
	sideEffect := func(int, int) {
		called++
	}

	var fn func(int, int) int
	mock.Wrap(&fn).           // wrap an existing function in a mock ...
		Expect(2, NotNil()).    // expect to be called with arguments 2 and any non-nil argument ...
		Times(NewAtMost(2)).    // either 0, 1 or 2 times ...
		SideEffect(sideEffect). // when we do, perform the side effect ...
		Returns(4)              // and return 4

	expect.Equals(4, fn(2, 3)) // matches
	expect.Equals(4, fn(2, 4)) // matches
	expect.Equals(0, fn(4, 5)) // does not match

	expect.Equals(2, called) // 2 of the above matched

	mock.Verify() // verification succeeds

	expect.Equals(4, fn(2, 5)) // call again

	mock.Verify() // verification failes, the AtMost(2) condition does not hold
}
```
