package runjs

import (
	"github.com/robertkrimen/otto"
)

type targets struct {
	strs []string
}

func (t *targets) add(fnc otto.FunctionCall) otto.Value {
	ot := fnc.Otto
	if len(fnc.ArgumentList) != 2 {
		return ot.MakeCustomError("PareError", "addTarget takes two arguments")
	}

	t.strs = append(t.strs, fnc.ArgumentList[0].String())
	return otto.TrueValue()
}
