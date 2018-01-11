package runjs

import (
	"github.com/robertkrimen/otto"
)

type jsTargets struct {
	funcs map[string]otto.Value
}

func newJSTargets() *jsTargets {
	return &jsTargets{funcs: make(map[string]otto.Value)}
}

func (t *jsTargets) add(fnc otto.FunctionCall) otto.Value {
	ot := fnc.Otto
	args := fnc.ArgumentList
	if len(args) != 2 {
		return newPareError(ot, "addTarget takes two arguments")
	}

	funcName, err := args[0].ToString()
	if err != nil {
		return newPareError(ot, "first argument to addTarget must be a string")
	}

	if !args[1].IsFunction() {
		return newPareError(ot, "second argument to addTarget must be a value")
	}
	t.funcs[funcName] = args[1]
	return otto.TrueValue()
}
