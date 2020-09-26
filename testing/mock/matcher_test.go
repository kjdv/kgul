package mock

import (
	"fmt"
	"testing"

	"github.com/kjdv/kgul/testing/expects"
)

func doMatch(m Matcher, v interface{}) bool {
	return m.Match(v)
}

func doString(s fmt.Stringer) string {
	return s.String()
}

func TestAny_Match(t *testing.T) {
	expect := expects.New(t)

	any := NewAny()
	expect.True(doMatch(any, 1))
	expect.True(doMatch(any, "blah"))
	expect.True(doMatch(any, nil))
}

func TestAnyT_String(t *testing.T) {
	expect := expects.New(t)

	any := NewAny()
	expect.Equals("Any()", doString(any))
}

func TestEquals_Match(t *testing.T) {
	expect := expects.New(t)

	eq := Eq(5)
	expect.True(doMatch(eq, 5))
	expect.False(doMatch(eq, 6))
	expect.False(doMatch(eq, nil))
}

func TestEquals_String(t *testing.T) {
	expect := expects.New(t)

	eq := Eq(4)
	expect.Equals("Eq(int=4)", doString(eq))
}

func TestNotEquals_Match(t *testing.T) {
	expect := expects.New(t)

	ne := Ne(5)
	expect.False(doMatch(ne, 5))
	expect.True(doMatch(ne, 6))
	expect.True(doMatch(ne, nil))
}

func TestNotEquals_String(t *testing.T) {
	expect := expects.New(t)

	ne := Ne(4)
	expect.Equals("Ne(int=4)", doString(ne))
}

func TestIsNilT_Match(t *testing.T) {
	expect := expects.New(t)

	in := IsNil()
	expect.False(doMatch(in, 9))
	expect.True(doMatch(in, nil))
}

func TestIsNilT_String(t *testing.T) {
	expect := expects.New(t)

	in := IsNil()
	expect.Equals("IsNil()", doString(in))
}

func TestNotNilT_Match(t *testing.T) {
	expect := expects.New(t)

	nn := NotNil()
	expect.True(doMatch(nn, 9))
	expect.False(doMatch(nn, nil))
}

func TestNotNilT_String(t *testing.T) {
	expect := expects.New(t)

	nn := NotNil()
	expect.Equals("NotNil()", doString(nn))
}

func TestAllOf_Match(t *testing.T) {
	expect := expects.New(t)

	ao := AllOf([]Matcher{Eq(5), Ne(6)})

	expect.True(doMatch(ao, []interface{}{5, 5}))
	expect.False(doMatch(ao, []interface{}{5}))
	expect.False(doMatch(ao, []interface{}{5, 5, 5}))
	expect.False(doMatch(ao, []interface{}{4, 5}))
	expect.False(doMatch(ao, []interface{}{5, 6}))
}

func TestAllOf_String(t *testing.T) {
	expect := expects.New(t)

	ao := AllOf([]Matcher{Eq(5), Ne(6)})

	expect.Equals("All Of([Eq(int=5) Ne(int=6)])", doString(ao))
}
