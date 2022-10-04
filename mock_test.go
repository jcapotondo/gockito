package gockito_test

import (
	"github.com/stretchr/testify/assert"
	"gockito"
	"testing"
)

type Asd interface {
	test() (error, int, *string)
	GetName(asd interface{}, a int) (string, int, error, interface{})
}

type testStruct struct {
	Name string
}

func (s testStruct) test() (error, int, *string) {
	return nil, 0, nil
}

func (s testStruct) GetName(asd interface{}, a int) (string, int, error, interface{}) {
	return "", 0, nil, nil
}

func TestNewMock(t *testing.T) {
	res := gockito.NewMock[Asd](t)

	res.Expect("GetName").With("few", false).Return("", 0, nil, nil)

	assert.Equal(t, 2, res)
}
