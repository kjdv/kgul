package mock

import (
	"testing"

	"github.com/kjdv/kgul/testing/expects"
)

func doCompare(c Cardinality, n int) (ok bool, msg string) {
	return c.Compare(n)
}

func TestNever_Compare(t *testing.T) {
	expect := expects.New(t)

	never := NewNever()

	ok, msg := doCompare(never, 0)
	expect.True(ok)
	expect.Equals("", msg)

	ok, msg = doCompare(never, 1)
	expect.False(ok)
	expect.NotEquals("", msg)
}

func TestNeverT_String(t *testing.T) {
	expect := expects.New(t)

	never := NewNever()
	expect.Equals("never called", doString(never))
}

func TestExactly_Compare(t *testing.T) {
	expect := expects.New(t)

	exact := NewExactly(3)

	ok, msg := doCompare(exact, 3)
	expect.True(ok)
	expect.Equals("", msg)

	ok, msg = doCompare(exact, 4)
	expect.False(ok)
	expect.NotEquals("", msg)

	ok, msg = doCompare(exact, 2)
	expect.False(ok)
	expect.NotEquals("", msg)
}

func TestExactly_String(t *testing.T) {
	expect := expects.New(t)

	exact := NewExactly(0)
	expect.Equals("called exactly 0 times", doString(exact))

	exact = NewExactly(1)
	expect.Equals("called exactly 1 time", doString(exact))

	exact = NewExactly(2)
	expect.Equals("called exactly 2 times", doString(exact))
}

func TestAtLeast_Compare(t *testing.T) {
	expect := expects.New(t)

	atleast := NewAtLeast(2)

	ok, msg := doCompare(atleast, 1)
	expect.False(ok)
	expect.NotEquals("", msg)

	ok, msg = doCompare(atleast, 2)
	expect.True(ok)
	expect.Equals("", msg)

	ok, msg = doCompare(atleast, 3)
	expect.True(ok)
	expect.Equals("", msg)
}

func TestAtLeastT_String(t *testing.T) {
	expect := expects.New(t)

	atleast := NewAtLeast(1)
	expect.Equals("called at least 1 time", doString(atleast))

	atleast = NewAtLeast(2)
	expect.Equals("called at least 2 times", doString(atleast))
}

func TestAtMost_Compare(t *testing.T) {
	expect := expects.New(t)

	atmost := NewAtMost(2)

	ok, msg := doCompare(atmost, 1)
	expect.True(ok)
	expect.Equals("", msg)

	ok, msg = doCompare(atmost, 2)
	expect.True(ok)
	expect.Equals("", msg)

	ok, msg = doCompare(atmost, 3)
	expect.False(ok)
	expect.NotEquals("", msg)
}

func TestAtMost_String(t *testing.T) {
	expect := expects.New(t)

	atmost := NewAtMost(1)
	expect.Equals("called at most 1 time", doString(atmost))

	atmost = NewAtMost(2)
	expect.Equals("called at most 2 times", doString(atmost))
}
