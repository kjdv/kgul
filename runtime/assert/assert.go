package assert

import (
	"fmt"
	"os"
	"runtime/debug"
)

type Failure struct {
	message string
	stack   []byte
}

func (f Failure) String() string {
	return f.message + "\n" + string(f.stack)
}

func DieHandler(f Failure) {
	fmt.Fprint(os.Stderr, f)
	os.Exit(1)
}

func PanicHandler(f Failure) {
	panic(f.message)
}

// Allow a custom function to be set to actually deal with assertion failures
// Default prints the message to stderr and exits
var assertHandler func(Failure) = DieHandler

func OnAssert(handler func(Failure)) {
	assertHandler = handler
}

func format(msg ...interface{}) string {
	var message string
	if len(msg) > 0 {
		message = fmt.Sprint(msg...)
	} else {
		message = "Assertion failed!"
	}

	return message
}

func formatf(format string, items ...interface{}) string {
	return fmt.Sprintf(format, items...)
}

func Assert(condition bool, msg ...interface{}) {
	if !condition {
		f := Failure{format(msg...), debug.Stack()}
		assertHandler(f)
	}
}

func Assertf(condition bool, format string, items ...interface{}) {
	if !condition {
		f := Failure{formatf(format, items...), debug.Stack()}
		assertHandler(f)
	}
}

func Fail(msg ...interface{}) {
	f := Failure{format(msg...), debug.Stack()}
	assertHandler(f)
}

func Failf(format string, items ...interface{}) {
	f := Failure{formatf(format, items...), debug.Stack()}
	assertHandler(f)
}

func NotErr(err error, msg ...interface{}) {
	Assert(err == nil, append([]interface{}{err}, msg...)...)
}

func NotErrf(err error, format string, items ...interface{}) {
	Assertf(err == nil, format, append([]interface{}{err}, items...)...)
}
