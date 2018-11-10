# kgul
Klaas' Go Utility library

A small set of assorted utilities to help with other Go projects.

## Assert

Makes it possible to use assertions, for example:

```go
import "github.com/klaasjacobdevries/kgul/runtime/assert"

func divide(a, b int) int {
	assert.Assert(b != 0, "Can not divide by zero")
	return a / b
}
```

This also comes in flavours, there is an `Assertf()` for formatted assertion messages in `Fail()` and `Failf()` to unconditionally fail, useful when flow-control rather than an explicit condition signals failure.

Default behaviour is to print the message and a stack trace to stderr and exit the process with a non-zero exit code. By calling `OnAssert()` a different failure handler can be installed:

```go
import "github.com/klaasjacobdevries/kgul/runtime/assert"

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
import "github.com/klaasjacobdevries/kgul/runtime/logging"

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

import "github.com/klaasjacobdevries/kgul/testing/expects"

func TestMyExpect(t *testing.T) {
	expect := expects.New(t)

	expect.True(2+2 == 5) // prints a friendly error message plus the file and line number of the failure
	expect.Equals(5, 2+2) // prints
	//  expected != actual
	//  expected: (struct { a string }) {a:foo}
	//  actual:   (struct { a string }) {a:bar}

	expect.Equals(struct{ a string }{"foo"}, struct{ a string }{"bar"} // complex structures supported, this fails with:
	//  expected != actual
	//  expected: (struct { a string }) {a:foo}
	//  actual:   (struct { a string }) {a:bar}
}
```

Other  expectations supported include `Fail()` (non-conditional, based on flow control), `IsNil()` and `IsNotNil()`, `AlmostEqual()` and `AlmostNotEqual()` (for floating-point comparisons),  `Less()`, `LessEqual()`, `Greater()` and `GreaterEqual()`, `Panics()`, `Regex()` for string matching and a generic `That()` to match for a custom matcher (c.f. Mock)
