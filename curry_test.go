package curry_test

import (
	"curry"
	"errors"
	"testing"
)

type AnyStruct struct {}

func (s AnyStruct) Sum(a, b int) int {
	return a + b
}

func crossMultiply(a, b, c float64) (float64, error) {
	if a == 0 {
		return 0, errors.New("can not divide by zero")
	}
	return (b * c) / a, nil
}

func TestFunc(t *testing.T) {
	var a, b, c float64 = 100, 420, 10
	cross, _ := curry.Func(crossMultiply)

	t.Run("partial apply works", func(t *testing.T) {
		partial, _ := cross(a)
		partial, _ = partial(b)
		_, out := partial(c)

		if out == nil || len(out) != 2 {
			t.Fail()
		}

		result := out[0].Float()
		err := out[1]

		if result != 42 || !err.IsNil() {
			t.Fail()
		}
	})
	t.Run("should be able to create new curried instances", func(t *testing.T) {
		var a float64 = 0

		partial, _ := cross(a)
		partial, _ = partial(b)
		_, out := partial(c)

		if out == nil || len(out) != 2 {
			t.Fail()
		}

		result := out[0].Float()
		err := out[1].Interface().(error)

		if result != 0 || err == nil || err.Error() != "can not divide by zero" {
			t.Fail()
		}
	})
}

func TestMethod(t *testing.T) {
	var a, b = 20, 22
	s := AnyStruct{}

	sum, _ := curry.Method(s, "Sum")
	partial, _ := sum(a)
	_, out := partial(b)

	if out == nil || len(out) != 1 {
		t.Fail()
	}

	result := out[0].Int()

	if result != 42 {
		t.Fail()
	}
}