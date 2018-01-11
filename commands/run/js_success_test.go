package run

import (
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
)

func TestSuccessObjetRoundTrip(t *testing.T) {
	const descr = "this is a test!"
	ot := otto.New()
	succVal, err := newSuccessObj(ot, &successObj{descrStr: descr})
	assert.NoError(t, err)
	retSuccObj, err := convertToSuccessObj(succVal)
	assert.NoError(t, err)
	assert.Equalf(t, &successObj{
		descrStr: descr,
	}, retSuccObj, "success objects didn't match")

}
