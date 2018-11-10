package expects

import (
	"fmt"
	"kdv/testing/metatest"
	"runtime"
)

func callpoint() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "???"
	}

	return fmt.Sprintf("%s:%d", file, line)
}

type Expect struct {
	t metatest.Tester
}

func New(t metatest.Tester) Expect {
	return Expect{t}
}
