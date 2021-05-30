package curry

import "reflect"

type Wrapper func(partialArgs ...interface{}) (Wrapper, []reflect.Value)

func Func(fn interface{}) (Wrapper, []reflect.Value) {
	return wrap(reflect.ValueOf(fn))
}

func Method(instance interface{}, method string) (Wrapper, []reflect.Value) {
	return wrap(reflect.ValueOf(instance).MethodByName(method))
}

func wrap(r reflect.Value) (Wrapper, []reflect.Value) {
	numIn := r.Type().NumIn()

	return func(partialArgs ...interface{}) (Wrapper, []reflect.Value) {
		var partial Wrapper

		args := make([]reflect.Value, numIn)
		idx := 0

		for _, arg := range partialArgs {
			args[idx] = reflect.ValueOf(arg)
			idx++
		}

		if idx == numIn {
			return nil, r.Call(args)
		}

		partial = func(partialArgs ...interface{}) (Wrapper, []reflect.Value) {
			for _, arg := range partialArgs {
				args[idx] = reflect.ValueOf(arg)
				idx++
			}

			if idx == numIn {
				return nil, r.Call(args)
			}

			return partial, nil
		}

		return partial, nil
	}, nil
}
