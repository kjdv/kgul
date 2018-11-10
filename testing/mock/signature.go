package mock

import (
	"fmt"
	"reflect"
)

func getFunction(fptr interface{}) (reflect.Value, error) {
	fn := reflect.ValueOf(fptr)
	if fn.Kind() != reflect.Func {
		fn = fn.Elem()
	}

	if fn.Type().Kind() != reflect.Func {
		return reflect.Value{}, fmt.Errorf("%t=%v is not a function or a pointer to a function", fptr, fptr)
	}

	return fn, nil
}

func asValues(args ...interface{}) []reflect.Value {
	values := []reflect.Value{}
	for _, arg := range args {
		values = append(values, reflect.ValueOf(arg))
	}

	return values
}

func asInterface(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.Invalid:
		return nil
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Ptr:
		fallthrough
	case reflect.Slice:
		if value.IsNil() {
			return nil
		}
	}

	return value.Interface()
}

type signature struct {
	typ reflect.Type
}

func deriveSignature(fptr interface{}) signature {
	fn, err := getFunction(fptr)
	if err != nil {
		panic(err)
	}
	return signature{fn.Type()}
}

func (s signature) makeReturnValue() []reflect.Value {
	returns := []reflect.Value{}
	for idx := 0; idx < s.typ.NumOut(); idx++ {
		returns = append(returns, reflect.New(s.typ.Out(idx)).Elem())
	}
	return returns
}

func (s signature) checkArguments(args ...interface{}) ([]reflect.Value, error) {
	if len(args) != s.typ.NumIn() {
		return nil, fmt.Errorf("expected %d arguments, but got %d", s.typ.NumIn(), len(args))
	}

	for idx, arg := range args {
		argt := reflect.TypeOf(arg)
		if !argt.ConvertibleTo(s.typ.In(idx)) {
			return nil, fmt.Errorf("argument #%d of type %v can not be converted to expected type %v", idx, argt, s.typ.In(idx))
		}
	}

	return asValues(args...), nil
}

func (s signature) checkReturns(rvs ...interface{}) ([]reflect.Value, error) {
	if len(rvs) != s.typ.NumOut() {
		return nil, fmt.Errorf("return value does not match signature: expected %d values, but got %d", s.typ.NumOut(), len(rvs))
	}

	for idx, rv := range rvs {
		rvt := reflect.TypeOf(rv)
		if !rvt.ConvertibleTo(s.typ.Out(idx)) {
			return nil, fmt.Errorf("value #%d of type %v can not be converted to expected type %v", idx, rvt, s.typ.Out(idx))
		}
	}

	return asValues(rvs...), nil
}

func (s signature) checkFunctionArguments(fptr interface{}) (reflect.Value, error) {
	fn, err := getFunction(fptr)
	zero := reflect.ValueOf(nil)

	if err != nil {
		return zero, err
	}

	fnt := fn.Type()

	if fnt.NumIn() != s.typ.NumIn() {
		return zero, fmt.Errorf("expected a function taking %d arguments, but got one taking %d", s.typ.NumIn(), fnt.NumIn())
	}

	for idx := 0; idx < fnt.NumIn(); idx++ {
		if !s.typ.In(idx).ConvertibleTo(fnt.In(idx)) {
			return zero, fmt.Errorf("cannot convert argument #%d of type %v to type %v", idx, s.typ.In(idx), fnt.In(idx))
		}
	}

	return fn, nil
}

func (s signature) isVariadic(argnum int) bool {
	return s.typ.IsVariadic() && argnum == s.typ.NumIn()-1
}

func (s signature) numArguments() (mandatory int, isVariadic bool) {
	if s.typ.IsVariadic() {
		return s.typ.NumIn() - 1, true
	}
	return s.typ.NumIn(), false
}
