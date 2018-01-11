package run

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/robertkrimen/otto"
)

type jsSuccess struct {
	target string
}

func (s *jsSuccess) run(fnc otto.FunctionCall) otto.Value {
	ot := fnc.Otto
	args := fnc.ArgumentList
	if len(args) != 1 && len(args) != 0 {
		return newPareError(ot, "success() takes zero or one arguments")
	}

	descrStr := ""
	if len(args) == 1 {
		descrVal := args[0]
		str, err := descrVal.ToString()
		if err != nil {
			return newPareError(ot, "first argument to success() was not a string")
		}
		descrStr = str
	}

	succObj, err := newSuccessObj(ot, &successObj{descrStr: descrStr})
	if err != nil {
		return newPareError(ot, "couldn't create a success object")
	}
	return succObj
}

type successObj struct {
	descrStr string
}

func (s *successObj) String() string {
	return s.descrStr
}

func newSuccessObj(ot *otto.Otto, succ *successObj) (otto.Value, error) {
	return ot.ToValue(map[string]string{"description": succ.descrStr})
}

func convertToSuccessObj(val otto.Value) (*successObj, error) {
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
	return &successObj{descrStr: descrStr}, nil
}
