package mock

import (
	"kdv/testing/metatest"
)

type Mock struct {
	t         metatest.Tester
	functions []*MockedFunction
}

func New(t metatest.Tester) Mock {
	return Mock{t, []*MockedFunction{}}
}

func (m *Mock) Wrap(fptr interface{}) *MockedFunction {
	mf := newMockedFunction(m.t, fptr)
	m.functions = append(m.functions, mf)
	return mf
}

func (m *Mock) Verify() {
	for _, f := range m.functions {
		f.Verify()
	}
}
