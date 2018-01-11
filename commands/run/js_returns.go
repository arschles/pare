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
	obj := val.Object()
	if obj == nil {
		return nil, fmt.Errorf("return value was not an object")
	}
	sRet, sErr := convertToSuccessReturn(obj)
	eRet, eErr := convertToErrorReturn(obj)
	if sErr != nil && eErr != nil {
		return nil, fmt.Errorf("unknown return value (%+v)", val)
	}
	if eErr == nil {
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
	return ot.ToValue(map[string]string{"pare_description": descr})
}

func convertToSuccessReturn(obj *otto.Object) (*successReturn, error) {
	descr, err := stringFromObject(obj, "pare_description")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &successReturn{descrStr: descr}, nil
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
		"pare_description": descr,
		"pare_exit_code":   int64(exitCode),
	})
}

func convertToErrorReturn(obj *otto.Object) (*errorReturn, error) {
	descr, err := stringFromObject(obj, "pare_description")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	code, err := numberFromObject(obj, "pare_exit_code")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &errorReturn{descrStr: descr, code: int(code)}, nil
}
