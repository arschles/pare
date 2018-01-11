package run

import (
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

	ret, err := newSuccessReturn(ot, descrStr)
	if err != nil {
		return newPareError(ot, "couldn't create a success object")
	}
	return ret
}
