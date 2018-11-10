package mock

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/klaasjacobdevries/kgul/testing/metatest"
)

func callpoint() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "???"
	}

	return fmt.Sprintf("%s:%d", file, line)
}

type MockedFunction struct {
	t    metatest.Tester
	sign signature

	calls [][]reflect.Value

	expectations []*Expectation
	callpoints   []string
}

func newMockedFunction(t metatest.Tester, fptr interface{}) *MockedFunction {
	fn := reflect.ValueOf(fptr).Elem()

	mf := MockedFunction{t, deriveSignature(fptr), [][]reflect.Value{}, []*Expectation{}, []string{}}

	v := reflect.MakeFunc(fn.Type(), mf.match)
	fn.Set(v)

	return &mf
}

func (mf *MockedFunction) match(args []reflect.Value) []reflect.Value {
	mf.calls = append(mf.calls, args)

	sargs := []string{}
	for _, arg := range args {
		i := asInterface(arg)
		sargs = append(sargs, fmt.Sprintf("%T=%v", i, i))
	}

	for idx, exp := range mf.expectations {
		rv, ok, msg := exp.match(args)
		if ok {
			return rv
		}

		cp := mf.callpoints[idx]
		mf.t.Logf("Not matching %v\nto %v\nbecause %v\nfrom %s", sargs, exp, msg, cp)
	}

	mf.t.Logf("Warning: no expectation for %v matched", sargs)
	return mf.sign.makeReturnValue()
}

func (mf *MockedFunction) NumCalls() int {
	return len(mf.calls)
}

func (mf *MockedFunction) CalledWith(arguments ...interface{}) bool {
	match := func(call []reflect.Value) bool {
		if len(arguments) == len(call) {
			for idx, _ := range call {
				if call[idx].Interface() != arguments[idx] {
					return false
				}
			}
			return true
		}
		return false
	}

	for _, c := range mf.calls {
		if match(c) {
			return true
		}
	}
	return false
}

func (mf *MockedFunction) Expect(arguments ...interface{}) *Expectation {
	matchers := []Matcher{}
	for _, argument := range arguments {
		matcher, ok := argument.(Matcher)
		if !ok {
			matcher = Eq(argument)
		}

		matchers = append(matchers, matcher)
	}

	if mandatory, is_variadic := mf.sign.numArguments(); is_variadic {
		mm := matchers[:mandatory]
		vm := AllOf(matchers[mandatory:])
		matchers = append(mm, vm)
	}

	exp := newExpectation(mf.sign, matchers)

	mf.expectations = append(mf.expectations, exp)
	mf.callpoints = append(mf.callpoints, callpoint())
	return exp
}

func (mf *MockedFunction) Verify() {
	for idx, exp := range mf.expectations {
		ok, msg := exp.verify()
		if !ok {
			cp := mf.callpoints[idx]
			mf.t.Errorf("%v\nnot met: %s\nfrom %s", exp, msg, cp)
		}
	}
}
