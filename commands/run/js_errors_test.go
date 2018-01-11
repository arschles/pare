package run

import (
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
)

func TestErrorObjectRoundTrip(t *testing.T) {
	const exitCode = 1
	const descr = "this is a test error"
	ot := otto.New()
	errVal, err := newPareErrorObject(ot, exitCode, descr)
	assert.NoError(t, err)
	retErrObj, err := convertToErrObj(errVal)
	assert.NoError(t, err)
	assert.Equalf(t, &errObj{
		descrStr: descr,
		code:     exitCode,
	}, retErrObj, "error objects didn't match")

}
