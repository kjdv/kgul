package mock

import (
	"testing"

	"github.com/klaasjacobdevries/kgul/testing/expects"
)

type variadicInvoker interface {
	invoke(a int, arg ...int) int
}

type mockVariadicInvoker func(int, ...int) int

func (m mockVariadicInvoker) invoke(a int, arg ...int) int {
	return m(a, arg...)
}

func TestVariadicExample(t *testing.T) {
	mock := New(t)
	defer mock.Verify()

	expect := expects.New(t)

	var w mockVariadicInvoker
	f := mock.Wrap(&w)

	f.Expect(4, 5, 6).Times(Exactly{1}).Returns(3)

	expect.Equals(3, w.invoke(4, 5, 6))
}
