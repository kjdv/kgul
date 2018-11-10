package mock

import (
	"testing"

	"github.com/klaasjacobdevries/kgul/testing/expects"
)

type invoker interface {
	invoke(a, b int) int
}

type mockInvoker func(int, int) int

func (m mockInvoker) invoke(a, b int) int {
	return m(a, b)
}

func TestExample(t *testing.T) {
	mock := New(t)
	defer mock.Verify()

	expect := expects.New(t)

	var w mockInvoker
	f := mock.Wrap(&w)

	f.Expect(4, 5).Times(Exactly{1}).Returns(3)

	expect.Equals(3, w.invoke(4, 5))
}
