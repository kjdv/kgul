package checks

import (
	"fmt"
)

// Checks are similar to Asserts, in the sense that they are meant as a defence for situations that your program intentionally does not handle.
//
// Checks differ from Asserts in the following ways:
//
// * Asserts are meant for pure programming errors, and should never fire on external input. Checks can fire on external input like a bad configuration file, or other situations
//   where panicking is considered appropriate.
// * Following from the above: asserts should always be optional. A program with asserts disabled should be functionally identical from one that has them enabled. Checks however are not optional.
// * Checks call panic. A failing assert will often kill the program, but in general can be anything (including nothing)

func format(msg ...interface{}) interface{} {
	if len(msg) == 1 {
		return msg[0]
	} else if len(msg) > 0 {
		return fmt.Sprint(msg...)
	} else {
		return "Check failed!"
	}
}

func Check(condition bool, msg ...interface{}) {
	if !condition {
		Fail(msg...)
	}
}

func Checkf(condition bool, format string, items ...interface{}) {
	if !condition {
		Failf(format, items...)
	}
}

func Fail(msg ...interface{}) {
	panic(format(msg...))
}

func Failf(format string, items ...interface{}) {
	panic(fmt.Sprintf(format, items...))
}

func NotErr(err error, msg ...interface{}) {
	Check(err == nil, append([]interface{}{err}, msg...)...)
}

func NotErrf(err error, format string, items ...interface{}) {
	Checkf(err == nil, format, append([]interface{}{err}, items...)...)
}
