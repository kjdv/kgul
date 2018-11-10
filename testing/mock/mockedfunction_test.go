package mock

import (
	"testing"

	"github.com/klaasjacobdevries/kgul/testing/expects"
	"github.com/klaasjacobdevries/kgul/testing/metatest"
)

func TestMockedFunction_NumberOfCalls(t *testing.T) {
	expect := expects.New(t)

	var fn func()
	mf := newMockedFunction(t, &fn)
	expect.Equals(0, mf.NumCalls())

	fn()
	expect.Equals(1, mf.NumCalls())
}

func TestMockedFunction_CalledWith(t *testing.T) {
	expect := expects.New(t)

	var fn func(a, b int) int
	mf := newMockedFunction(t, &fn)

	fn(2, 3)
	fn(3, 4)

	expect.True(mf.CalledWith(2, 3))
	expect.True(mf.CalledWith(3, 4))
	expect.False(mf.CalledWith(4, 5))
	expect.False(mf.CalledWith(2, 3, 4))
}

func TestMockedFunction_Expect(t *testing.T) {
	mt := metatest.New()
	expect := expects.New(t)

	var fn func(a, b int) int
	mf := newMockedFunction(mt, &fn)
	mf.Expect(1, 2).NTimes(2)

	fn(1, 2)
	fn(1, 3)
	fn(1, 2)

	mf.Verify()
	expect.False(mt.HasErrors(), mt)

	fn(1, 2)
	mf.Verify()
	expect.True(mt.HasErrors(), mt)
}

func TestMockedFunction_Variadic(t *testing.T) {
	mt := metatest.New()
	expect := expects.New(t)

	var fn func(int, ...int)
	mf := newMockedFunction(mt, &fn)

	expectation := mf.Expect(1)
	expect.Equals(2, len(expectation.matchers), "should be 1 for the mandatory argument and a slice for the variadic one")

	expectation = mf.Expect(1, 2)
	expect.Equals(2, len(expectation.matchers), "should be 1 for the mandatory argument and a slice for the variadic one")

	expectation = mf.Expect(1, 2, 3)
	expect.Equals(2, len(expectation.matchers), "should be 1 for the mandatory argument and a slice for the variadic one")
}
