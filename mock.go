package gockito

import "reflect"

type method struct {
	name       string
	paramsKind []reflect.Kind
}

type Mock struct {
	methods []method
}

func NewMock[T any](str T) Mock {
	strType := reflect.TypeOf(str)

	switch strType.Kind() {
	case reflect.Struct:
		return mockStruct(strType)
	}

	return Mock{}
}

func mockStruct(strType reflect.Type) Mock {
	methods := make([]method, 0)
	for i := 0; i < strType.NumMethod(); i++ {
		m := strType.Method(i)

		value := method{
			name:       m.Name,
			paramsKind: getStructParametersKind(m),
		}

		methods = append(methods, value)
	}

	return Mock{
		methods: methods,
	}
}

func getStructParametersKind(m reflect.Method) []reflect.Kind {
	fnType := m.Type

	res := make([]reflect.Kind, 0)

	// numOfParameters is the amount of parameters that the function has.
	// When its equals to 1 is because there are no parameters.
	numOfParameters := fnType.NumIn()

	for i := 1; i < numOfParameters; i++ {
		paramType := fnType.In(i)

		res = append(res, paramType.Kind())
	}

	return res
}
