package mock

import (
	"fmt"
	"reflect"
)

type Matcher interface {
	Match(value interface{}) bool
}

type Any struct {
}

func NewAny() Any {
	return Any{}
}

func (a Any) Match(value interface{}) bool {
	return true
}

func (a Any) String() string {
	return "Any()"
}

type Equals struct {
	value interface{}
}

func Eq(value interface{}) Equals {
	return Equals{value}
}

func (e Equals) Match(value interface{}) bool {
	return e.value == value
}

func (e Equals) String() string {
	return fmt.Sprintf("Eq(%T=%v)", e.value, e.value)
}

type NotEquals struct {
	value interface{}
}

func Ne(value interface{}) NotEquals {
	return NotEquals{value}
}

func (n NotEquals) Match(value interface{}) bool {
	return n.value != value
}

func (n NotEquals) String() string {
	return fmt.Sprintf("Ne(%T=%v)", n.value, n.value)
}

type IsNilT struct {
}

func IsNil() IsNilT {
	return IsNilT{}
}

func (in IsNilT) Match(value interface{}) bool {
	return value == nil
}

func (in IsNilT) String() string {
	return "IsNil()"
}

type NotNilT struct {
}

func NotNil() NotNilT {
	return NotNilT{}
}

func (nn NotNilT) Match(value interface{}) bool {
	return value != nil
}

func (nn NotNilT) String() string {
	return "NotNil()"
}

type AllOfT struct {
	matchers []Matcher
}

func AllOf(matchers []Matcher) AllOfT {
	ao := AllOfT{make([]Matcher, len(matchers))}
	copy(ao.matchers, matchers)
	return ao
}

func (ao AllOfT) Match(value interface{}) bool {
	values := reflect.ValueOf(value)

	if values.Len() != len(ao.matchers) {
		return false
	}

	for idx := 0; idx < values.Len(); idx++ {
		m := ao.matchers[idx]

		if !m.Match(values.Index(idx).Interface()) {
			return false
		}
	}

	return true
}

func (ao AllOfT) String() string {
	return fmt.Sprintf("All Of(%v)", ao.matchers)
}
