package runjs

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no target was specified")
	}

	fileBytes, err := ioutil.ReadFile("pare.js")
	if err != nil {
		logger.Printf("error reading file")
		return errors.New("error reading file")
	}
	vm := otto.New()
	targets := &targets{}
	vm.Set("addTarget", targets.add)
	val, err := vm.Run(string(fileBytes))
	if err != nil {
		logger.Printf("error running (%s)", err)
		return err
	}
	logger.Printf("found targets (%+v)", targets.strs)
	logger.Printf("success! (%+v)", val)
	return nil
}
