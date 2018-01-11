package run

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/robertkrimen/otto"
)

type jsReturnObject interface {
	fmt.Stringer
}

func convertToReturnObj(ot *otto.Otto, val otto.Value) (jsReturnObject, error) {
	sRet, sErr := convertToSuccessReturn(val)
	eRet, eErr := convertToErrorReturn(val)
	if sErr != nil && eErr != nil {
		return nil, fmt.Errorf("unknown return value (%+v)", val)
	}
	if sErr != nil {
		return eRet, nil
	}
	return sRet, nil

}

type successReturn struct {
	descrStr string
}

func (s *successReturn) String() string {
	return s.descrStr
}

func newSuccessReturn(ot *otto.Otto, descr string) (otto.Value, error) {
	return ot.ToValue(map[string]string{"description": descr})
}

func convertToSuccessReturn(val otto.Value) (*successReturn, error) {
	obj := val.Object()
	if obj == nil {
		return nil, errors.WithStack(fmt.Errorf("no success object exists"))
	}
	descrVal, err := obj.Get("description")
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("no description found"))
	}
	descrStr, err := descrVal.ToString()
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("description was not a string"))
	}
	return &successReturn{descrStr: descrStr}, nil
}

type errorReturn struct {
	descrStr string
	code     int
}

func (e *errorReturn) Error() string {
	return fmt.Sprintf("%s (exit code %d)", e.descrStr, e.code)
}

func (e *errorReturn) String() string {
	return fmt.Sprintf("%s (exit code %d)", e.descrStr, e.code)
}

func newErrorReturn(ot *otto.Otto, exitCode int, descr string) (otto.Value, error) {
	return ot.ToValue(map[string]interface{}{
		"error": descr,
		"code":  exitCode,
	})
}

func convertToErrorReturn(val otto.Value) (*errorReturn, error) {
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
	return &errorReturn{descrStr: descrStr, code: int(errCode)}, nil
}
