package run

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/robertkrimen/otto"
)

func callJSFunction(ot *otto.Otto, targetName string, val otto.Value) error {
	callVal, err := val.Call(val)
	if err != nil {
		color.Red("-----> error calling target %s (%s)", targetName, err)
		return fmt.Errorf("error calling target %s (%s)", targetName, err)
	}
	retObj, err := convertToReturnObj(ot, callVal)
	if err != nil {
		return errors.WithStack(err)
	}
	switch t := retObj.(type) {
	case *successReturn:
		color.Green(t.descrStr)
		return nil
	case *errorReturn:
		color.Red("Error: %s", t.descrStr)
		os.Exit(t.code) // TODO: return error indicating the exit code
	}
	return fmt.Errorf("unknown return from script (%+v)", callVal)
}
