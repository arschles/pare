package run

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/robertkrimen/otto"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no target was specified")
	}
	targetName := args[0]

	vm := otto.New()
	targets := newJSTargets()
	vm.Set("addTarget", targets.add)
	vm.Set("error", jsErrorFunc)
	success := &jsSuccess{target: targetName}
	vm.Set("success", success.run)
	vm.Set("cmd", jsCmd)
	script, err := vm.Compile("pare.js", nil)
	if err != nil {
		color.Red("-----> error compiling script (%s)", err)
		return fmt.Errorf("error compiling script (%s)", err)
	}

	color.Green("----> successfully compiled js file")

	val, err := vm.Run(script)
	if err != nil {
		logger.Printf("error running (%s)", err)
		return err
	}

	val, ok := targets.funcs[targetName]
	if !ok {
		color.Red("-----> target %s not found", targetName)
		return fmt.Errorf("target %s not found", targetName)
	}
	return callJSFunction(vm, targetName, val)
}
