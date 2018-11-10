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

Default behaviour is to print the message to stderr and exit the process with a non-zero exit code. By calling `OnAssert()` a different failure handler can be installed:

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
