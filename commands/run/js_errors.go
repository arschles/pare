package run

import (
	"github.com/robertkrimen/otto"
)

func newPareError(ot *otto.Otto, descr string) otto.Value {
	return ot.MakeCustomError("PareError", descr)
}
