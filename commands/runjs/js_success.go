package runjs

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

	if len(args) == 1 {
		descrVal := args[0]
		str, err := descrVal.ToString()
		if err != nil {
			return newPareError(ot, "first argument to success() was not a string")
		}
		logger.Printf(str)
	}

	return otto.UndefinedValue()
}
