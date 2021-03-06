package run

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func jsErrorFunc(fnc otto.FunctionCall) otto.Value {
	ot := fnc.Otto
	args := fnc.ArgumentList
	if len(args) != 2 {
		return newPareError(ot, "error() takes two arguments")
	}

	exitVal := args[0]
	descrVal := args[1]

	if !exitVal.IsNumber() {
		return newPareError(ot, "first argument to error() is not a number")
	}
	if !descrVal.IsString() {
		return newPareError(ot, "second argument to error() is not a string")
	}

	exitNum, err := exitVal.ToInteger()
	if err != nil {
		return newPareError(ot, "first argument to error() is not a number")
	}
	descrStr, err := descrVal.ToString()
	if err != nil {
		return newPareError(ot, "second argument to error() is not a string")
	}
	ret, err := newErrorReturn(ot, int(exitNum), descrStr)
	if err != nil {
		return newPareError(ot, fmt.Sprintf("couldn't create a new error object (%s)", err))
	}
	return ret
}
