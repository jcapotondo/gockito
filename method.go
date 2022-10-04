package gockito

import (
	"reflect"
	"testing"
)

type Method struct {
	t                 *testing.T
	name              string
	expectedInParams  []userParameter
	expectedOutParams []userParameter
	paramsIn          []parameter
	paramsOut         []parameter
}

func newMethod(t *testing.T, m reflect.Method) Method {
	return Method{
		t:         t,
		name:      m.Name,
		paramsIn:  getInParameters(m),
		paramsOut: getOutParameters(m),
	}
}

func (m Method) With(parameters ...interface{}) Method {
	if len(m.paramsIn) != len(parameters) {
		m.t.Logf("expected `%d` parameters and got `%d`", len(m.paramsIn), len(parameters))
		m.t.FailNow()
	}

	m.expectedInParams = getUserParameters(parameters)

	return m
}

func (m Method) Return(parameters ...interface{}) {
	if len(m.paramsOut) != len(parameters) {
		m.t.Logf("expected `%d` parameters to be returned and got `%d`", len(m.paramsOut), len(parameters))
		m.t.FailNow()
	}

	m.expectedOutParams = getUserParameters(parameters)
}
