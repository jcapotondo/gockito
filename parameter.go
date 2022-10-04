package gockito

import "reflect"

type parameter struct {
	// Name of the type
	kindName string

	// Value in reflection of the type of the parameter
	kind reflect.Kind
}

// getInParameters returns the parameters that the function expects to receive.
func getInParameters(m reflect.Method) []parameter {
	fnType := m.Type

	return getParameters(fnType.NumIn(), fnType.In)
}

// getInParameters returns the parameters that the function returns.
func getOutParameters(m reflect.Method) []parameter {
	fnType := m.Type

	return getParameters(fnType.NumOut(), fnType.Out)
}

func getParameters(n int, fn func(j int) reflect.Type) []parameter {
	res := make([]parameter, 0)

	for i := 0; i < n; i++ {
		paramType := fn(i)

		param := parameter{
			kindName: paramType.String(),
			kind:     paramType.Kind(),
		}

		res = append(res, param)
	}

	return res
}

type userParameter struct {
	// Name of the type
	kindName string

	// Value in reflection of the type of the parameter
	kind reflect.Kind

	value interface{}
}

func getUserParameters(parameters ...interface{}) []userParameter {
	res := make([]userParameter, 0)
	for _, v := range parameters {
		paramType := reflect.TypeOf(v)

		res = append(res, userParameter{
			kindName: paramType.String(),
			kind:     paramType.Kind(),
			value:    v,
		})
	}

	return res
}
