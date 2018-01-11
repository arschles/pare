package run

import (
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
)

func TestConvertToErrorReturn(t *testing.T) {
	const exitCode = 1
	const descr = "this is a test error"
	ot := otto.New()
	errVal, err := newErrorReturn(ot, exitCode, descr)
	assert.NoError(t, err)
	retObj, err := convertToReturnObj(ot, errVal)
	assert.NoError(t, err)
	retError, ok := retObj.(*errorReturn)
	assert.True(t, ok, "expected an error return, got (%#v)", retObj)
	assert.Equal(t, exitCode, retError.code, "exit codes didn't match")
	assert.Equal(t, descr, retError.descrStr, "description strings didn't match")
}

func TestConvertToSuccessReturn(t *testing.T) {
	const msg = "test success!"
	vm := otto.New()
	successRet, err := newSuccessReturn(vm, msg)
	assert.NoError(t, err)
	retObj, err := convertToReturnObj(vm, successRet)
	assert.NoError(t, err)
	retSuccess, ok := retObj.(*successReturn)
	assert.True(t, ok, "expected a success return, got %#v", retObj)
	assert.Equal(t, msg, retSuccess.descrStr, "description strings didn't match")
}
