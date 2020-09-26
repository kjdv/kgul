package mock

import (
	"testing"

	"github.com/kjdv/kgul/testing/expects"
	"github.com/kjdv/kgul/testing/metatest"
)

func someFunc(callback func(int) int, arg int) int {
	return callback(arg)
}

func TestMock_Wrap(t *testing.T) {
	mock := New(t)
	expect := expects.New(t)

	var w func(int) int
	f := mock.Wrap(&w)
	expect.Equals(0, someFunc(w, 3))
	expect.Equals(0, someFunc(w, 5))

	expect.Equals(2, f.NumCalls())
}

type someInterface interface {
	doSomething() int
}

func takesInterface(i someInterface) int {
	return i.doSomething()
}

type mockInterfacer func() int

func (m mockInterfacer) doSomething() int {
	return m()
}

func TestMock_WrapInterface(t *testing.T) {
	m := New(t)
	expect := expects.New(t)

	var w mockInterfacer
	f := m.Wrap(&w)
	expect.Equals(0, takesInterface(w))

	expect.Equals(1, f.NumCalls())
}

func TestMock_Verify(t *testing.T) {
	mt := metatest.New()
	expect := expects.New(t)
	mock := New(mt)

	called := 0
	sideEffect := func(int, int) {
		called++
	}

	var fn func(int, int) int
	mock.Wrap(&fn).
		Expect(2, NotNil()).
		Times(NewAtMost(2)).
		SideEffect(sideEffect).
		Returns(4)

	expect.Equals(4, fn(2, 3))
	expect.Equals(4, fn(2, 4))
	expect.Equals(0, fn(4, 5))

	expect.Equals(2, called)

	mock.Verify()
	expect.False(mt.HasErrors(), mt)

	expect.Equals(4, fn(2, 5))

	mock.Verify()
	expect.True(mt.HasErrors(), mt)
}

func TestMock_Verify_interface(t *testing.T) {
	mt := metatest.New()
	expect := expects.New(t)
	mock := New(mt)

	var w mockInterfacer
	mock.Wrap(&w).
		Expect().
		Times(NewAtMost(2)).
		Returns(5)

	expect.Equals(5, takesInterface(w))
	expect.Equals(5, takesInterface(w))

	mock.Verify()
	expect.False(mt.HasErrors(), mt)

	expect.Equals(5, takesInterface(w))
	mock.Verify()
	expect.True(mt.HasErrors(), mt)
}
