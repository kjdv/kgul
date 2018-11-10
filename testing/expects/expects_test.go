package expects

import (
	"testing"

	"github.com/klaasjacobdevries/kgul/testing/metatest"
)

func shouldSucceed(mt *metatest.Metatest, t *testing.T) {
	if mt.HasErrors() {
		t.Error("No errors expected:", mt)
	}

	mt.Clear()
}

func shouldFail(mt *metatest.Metatest, t *testing.T) {
	if !mt.HasErrors() {
		t.Error("Expected to fail:", mt)
	}

	mt.Clear()
}
