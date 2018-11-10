package metatest

import "testing"

func TestMetatest_Errors(t *testing.T) {
	mt := New()

	if actual := mt.String(); actual != "no errors" {
		t.Error("unexpected, got '" + actual + "'")
	}
	if mt.HasErrors() {
		t.Error("expected no errors")
	}
	if num := mt.NumErrors(); num != 0 {
		t.Error("expected num errors to be 0")
	}
	if _, err := mt.LastError(); err == nil {
		t.Error("err expected to be not-nil")
	}

	mt.Error("some ", 3, " arguments")
	if actual := mt.String(); actual != "some 3 arguments" {
		t.Error("unexpected, got '" + actual + "'")
	}
	if !mt.HasErrors() {
		t.Error("errors expected")
	}
	if actual, err := mt.LastError(); !(err == nil) || actual != "some 3 arguments" {
		t.Error("unexpected, got '" + actual + "'")
	}
	if actual := mt.GetError(0); actual != "some 3 arguments" {
		t.Error("unexpected, got '" + actual + "'")
	}

	mt.Errorf("some %d format string", 4)
	if actual := mt.String(); actual != "some 3 arguments,\nsome 4 format string" {
		t.Error("unexpected, got '" + actual + "'")
	}

	if num := mt.NumErrors(); num != 2 {
		t.Error("expected num errors to be 2")
	}
	if actual, err := mt.LastError(); !(err == nil) || actual != "some 4 format string" {
		t.Error("unexpected, got '" + actual + "'")
	}
	if actual := mt.GetError(0); actual != "some 3 arguments" {
		t.Error("unexpected, got '" + actual + "'")
	}

	mt.Clear()

	if actual := mt.String(); actual != "no errors" {
		t.Error("unexpected, got '" + actual + "'")
	}
	if mt.HasErrors() {
		t.Error("expected no errors")
	}
	if num := mt.NumErrors(); num != 0 {
		t.Error("expected num errors to be 0")
	}
	if _, err := mt.LastError(); err == nil {
		t.Error("err expected to be not-nil")
	}
}

func TestMetatest_Logs(t *testing.T) {
	mt := New()

	if len(mt.logs) > 0 {
		t.Error("expected to be empty")
	}

	mt.Log("some ", "items")
	if len(mt.logs) != 1 || mt.logs[0] != "some items" {
		t.Fail()
	}

	mt.Logf("format %d string", 2)
	if len(mt.logs) != 2 || mt.logs[1] != "format 2 string" {
		t.Fail()
	}

	mt.Clear()
	if len(mt.logs) > 0 {
		t.Error("expected to be empty")
	}
}
