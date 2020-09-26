package expects

import (
	"testing"

	"github.com/kjdv/kgul/testing/metatest"
)

type isOdd struct{}

func (m isOdd) Match(value interface{}) bool {
	v := value.(int)

	return v&1 != 0
}

func (m isOdd) String() string {
	return "Matcher isOdd"
}

func TestExpect_That(t *testing.T) {
	mt := metatest.New()
	e := New(mt)

	m := isOdd{}
	e.That(3, m, "this matches")
	shouldSucceed(mt, t)

	e.That(4, m, "this does not match")
	shouldFail(mt, t)
}
