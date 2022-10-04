package gockito

import (
	"reflect"
	"testing"
)

type Mock struct {
	t          *testing.T
	structName string
	methods    map[string]Method
}

func (m Mock) Expect(name string) Method {
	method, exists := m.methods[name]
	if !exists {
		m.t.Logf("Unexpected call, `%s` does not exists in `%s`", name, m.structName)
		m.t.FailNow()
	}

	return method
}

type RES struct {
}

func NewMock[T any](t *testing.T) Mock {
	instance := new(T)

	element := reflect.TypeOf(instance).Elem()

	switch element.Kind() {
	case reflect.Interface:
		return generateMock(t, element)
	default:
		t.Log("Provided type MUST be an interface")
		t.FailNow()
		return Mock{}
	}
}

func generateMock(t *testing.T, strType reflect.Type) Mock {
	methods := make(map[string]Method)

	for i := 0; i < strType.NumMethod(); i++ {
		m := strType.Method(i)

		methods[m.Name] = newMethod(t, m)
	}

	return Mock{
		t:          t,
		structName: strType.String(),
		methods:    methods,
	}
}
