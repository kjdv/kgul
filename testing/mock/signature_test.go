package mock

import (
	"testing"

	"github.com/kjdv/kgul/testing/expects"
)

func TestValueConversion(t *testing.T) {
	expect := expects.New(t)

	convert := func(arg interface{}) interface{} {
		vs := asValues(arg)
		expect.Equals(1, len(vs))
		return asInterface(vs[0])
	}

	expect.Equals(1, convert(1))

	ch := make(chan int)
	expect.Equals(ch, convert(ch))

	// fn := func(){}
	// expect.Equals(fn, convert(fn)) // uncomparable type

	var ifc interface{} = 5
	expect.Equals(5, convert(ifc))

	mp := map[string]string{
		"foo": "bar",
	}
	expect.Equals("bar", (convert(mp)).(map[string]string)["foo"])

	mptr := &mp
	expect.Equals(mptr, convert(mptr))

	sl := []int{1, 2, 3}
	csl := convert(sl).([]int)
	expect.Equals(1, csl[0])

	sl = nil
	expect.IsNil(convert(sl))
}

func TestSignature_deriveNonFunction(t *testing.T) {
	expect := expects.New(t)

	s := "blah"
	expect.Panics(func() { deriveSignature(s) })
	expect.Panics(func() { deriveSignature(&s) })
}

func TestSignature_numberOfArguments(t *testing.T) {
	expect := expects.New(t)

	fn := func(a, b int) {}
	sig := deriveSignature(&fn)

	_, err := sig.checkArguments(1, 2)
	expect.IsNil(err)

	_, err = sig.checkArguments(1)
	expect.NotNil(err)

	_, err = sig.checkArguments(1, 2, 3)
	expect.NotNil(err)
}

func TestSignature_typeOfArguments(t *testing.T) {
	expect := expects.New(t)

	fn := func(a int, b float64) {}
	sig := deriveSignature(&fn)

	_, err := sig.checkArguments(1, 2.0)
	expect.IsNil(err)

	_, err = sig.checkArguments("1", 2.0)
	expect.NotNil(err)

	_, err = sig.checkArguments(1, "2.0")
	expect.NotNil(err)

	_, err = sig.checkArguments(1, 2)
	expect.IsNil(err)
}

func TestSignature_numberOfReturnValues(t *testing.T) {
	expect := expects.New(t)

	fn := func() (int, int) { return 0, 0 }
	sig := deriveSignature(&fn)

	_, err := sig.checkReturns(1, 2)
	expect.IsNil(err)

	_, err = sig.checkReturns(1)
	expect.NotNil(err)

	_, err = sig.checkReturns(1, 2, 3)
	expect.NotNil(err)
}

func TestSignature_typeOfReturnValues(t *testing.T) {
	expect := expects.New(t)

	fn := func() (int, float64) { return 1, 2.0 }
	sig := deriveSignature(&fn)

	_, err := sig.checkReturns(1, 2.0)
	expect.IsNil(err)

	_, err = sig.checkReturns("1", 2.0)
	expect.NotNil(err)

	_, err = sig.checkReturns(1, "2.0")
	expect.NotNil(err)
}

func TestSignature_checkFunctionArguments(t *testing.T) {
	expect := expects.New(t)

	fn := func(a int, b float64) {}
	sig := deriveSignature(&fn)

	good := func(a int, b float64) {}
	_, err := sig.checkFunctionArguments(&good)
	expect.IsNil(err)

	wrongNumber := func(a int, b float64, c string) {}
	_, err = sig.checkFunctionArguments(wrongNumber)
	expect.NotNil(err)

	wrongType := func(a int, b string) {}
	_, err = sig.checkFunctionArguments(&wrongType)
	expect.NotNil(err)

	notAFunction := "booh"
	_, err = sig.checkFunctionArguments(&notAFunction)
	expect.NotNil(err)
}

type sampleStruct struct {
}

func (s sampleStruct) method(a, b int) {
}

func TestSignature_checkFunctionArgumentsOfMethod(t *testing.T) {
	expect := expects.New(t)

	fn := func(a int, b int) {}
	sig := deriveSignature(&fn)

	str := sampleStruct{}
	_, err := sig.checkFunctionArguments(str.method)
	expect.IsNil(err)
}

func TestSignature_isVariadic(t *testing.T) {
	expect := expects.New(t)

	nvf := func(int, int) {}
	vf := func(int, ...int) {}
	non_variadic := deriveSignature(&nvf)
	variadic := deriveSignature(&vf)

	n, v := non_variadic.numArguments()
	expect.Equals(2, n)
	expect.False(v)

	n, v = variadic.numArguments()
	expect.Equals(1, n)
	expect.True(v)

	expect.False(variadic.isVariadic(0))
	expect.True(variadic.isVariadic(1))
	expect.False(non_variadic.isVariadic(1))
}
