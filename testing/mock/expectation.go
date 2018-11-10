package mock

import (
	"fmt"
	"reflect"
)

type Expectation struct {
	sign signature

	matchers    []Matcher
	cardinality Cardinality
	numMatches  int

	sideEffects []reflect.Value
	returnValue []reflect.Value
}

func newExpectation(s signature, m []Matcher) *Expectation {
	return &Expectation{
		s,
		m,
		NewAtLeast(1),
		0,
		[]reflect.Value{},
		s.makeReturnValue(),
	}
}

func (e *Expectation) String() string {
	return fmt.Sprint("expectation with arguments ", e.matchers, "\nto be ", e.cardinality)
}

func (e *Expectation) NumMatches() int {
	return e.numMatches
}

func (e *Expectation) NTimes(times int) *Expectation {
	if times == 0 {
		e.cardinality = Never{}
	} else {
		e.cardinality = Exactly{times}
	}

	return e
}

func (e *Expectation) Times(card Cardinality) *Expectation {
	e.cardinality = card
	return e
}

func (e *Expectation) SideEffect(fptr interface{}) *Expectation {
	fn, err := e.sign.checkFunctionArguments(fptr)
	if err != nil {
		panic(err)
	}

	e.sideEffects = append(e.sideEffects, fn)
	return e
}

func (e *Expectation) Returns(rv ...interface{}) *Expectation {
	values, err := e.sign.checkReturns(rv...)
	if err != nil {
		panic(err)
	}

	e.returnValue = values
	return e
}

func (e *Expectation) match(arguments []reflect.Value) (rv []reflect.Value, ok bool, msg string) {
	if len(arguments) != len(e.matchers) {
		return nil, false, fmt.Sprintf("wrong number of arguments: expected %d but got %d", len(e.matchers), len(arguments))
	}

	for idx, matcher := range e.matchers {
		arg := arguments[idx]
		if !matcher.Match(asInterface(arg)) {
			return nil, false, fmt.Sprintf("argument #%d (%v) did not match %s", idx, arg, matcher)
		}
	}

	e.numMatches++

	for _, se := range e.sideEffects {
		se.Call(arguments)
	}

	return e.returnValue, true, ""
}

func (e *Expectation) verify() (ok bool, msg string) {
	return e.cardinality.Compare(e.numMatches)
}
