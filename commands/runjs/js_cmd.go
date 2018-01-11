package runjs

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/robertkrimen/otto"
)

func jsCmd(fnc otto.FunctionCall) otto.Value {
	ot := fnc.Otto
	args := fnc.ArgumentList
	if len(args) == 0 {
		return newPareError(ot, "cmd() requires at least one command")
	}
	argStrings := make([]string, len(args))
	for i, arg := range args {
		s, err := arg.ToString()
		if err != nil {
			return newPareError(ot, fmt.Sprintf("argument %d must be a string", i))
		}
		argStrings[i] = s
	}

	cmdJoined := strings.Join(argStrings, " ")
	cmd := exec.Command(argStrings[0], argStrings[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return newPareError(ot, fmt.Sprintf("'%s', failed (%s)", cmdJoined, err))
	}
	retNum, err := ot.ToValue(0)
	if err != nil {
		return newPareError(ot, "command succeeded but couldn't convert exit code")
	}
	return retNum
}
