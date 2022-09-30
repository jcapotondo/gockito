package gockito_test

import (
	"github.com/stretchr/testify/assert"
	"gockito"
	"testing"
)

type testStruct struct {
	Name string
}

func (s testStruct) GREGREGER() {

}

func (s testStruct) GetName(asd string, a int) {

}

func TestNewMock(t *testing.T) {
	res := gockito.NewMock[testStruct](testStruct{})

	assert.Equal(t, 2, res)
}
