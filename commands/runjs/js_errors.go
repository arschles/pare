package runjs

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func newPareError(ot *otto.Otto, descr string) otto.Value {
	return ot.MakeCustomError("PareError", descr)
}

type errObj struct {
	descrStr string
	code     int
}

func (e *errObj) Error() string {
	return fmt.Sprintf("%s (exit code %d)", e.descrStr, e.code)
}

func newPareErrorObject(ot *otto.Otto, exitCode int, descr string) (otto.Value, error) {
	return ot.ToValue(map[string]interface{}{
		"error": descr,
		"code":  exitCode,
	})
}

func convertToErrObj(val otto.Value) (*errObj, error) {
	obj := val.Object()
	if obj == nil {
		return nil, fmt.Errorf("no error object exists")
	}
	descrVal, err := obj.Get("error")
	if err != nil {
		return nil, fmt.Errorf("couldn't find the error key")
	}

	codeVal, err := obj.Get("code")
	if err != nil {
		return nil, fmt.Errorf("couldn't find the code key")
	}

	descrStr, err := descrVal.ToString()
	if err != nil {
		return nil, fmt.Errorf("error description wasn't a string")
	}
	errCode, err := codeVal.ToInteger()
	if err != nil {
		return nil, fmt.Errorf("error code wasn't a number")
	}
	return &errObj{descrStr: descrStr, code: int(errCode)}, nil
}
