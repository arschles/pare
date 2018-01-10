package run

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/arschles/pare/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no target was specified")
	}
	cfg := new(config.File)
	/*tomlMD*/ _, err := toml.DecodeFile("pare.toml", cfg)
	if err != nil {
		return fmt.Errorf("Could not decode configuration file (%s)", err)
	}

	targetName := args[0]
	target, ok := cfg.Targets[targetName]
	if !ok {
		return fmt.Errorf("no target '%s' was in the config file", targetName)
	}

	startCtx, startFunc := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	resChan := make(chan cmdResult)
	for cmdName, cmd := range target.Commands {
		logger.Printf("-----> %s", cmd.Exec)
		wg.Add(1)
		c := make(chan cmdResult)
		go func(cmd *config.Command, resCh chan<- cmdResult) {
			<-startCtx.Done()
			defer close(resCh)
			if cmd.Exec == "" {
				return
			}
			spl := strings.Split(cmd.Exec, " ")
			first := spl[0]
			rest := spl[1:]
			runnable := exec.Command(first, rest...)
			runnable.Dir = cmd.Directory
			stdoutBuf := new(bytes.Buffer)
			stderrBuf := new(bytes.Buffer)
			runnable.Stdout = stdoutBuf
			runnable.Stderr = stderrBuf
			err := runnable.Run()
			resCh <- cmdResult{
				cmdName: cmdName,
				cfgCmd:  cmd,
				stdout:  stdoutBuf,
				stderr:  stderrBuf,
				err:     err,
				crash:   cmd.Crash,
			}
		}(cmd, c)
		go func(ch <-chan cmdResult) {
			defer wg.Done()
			res := <-ch
			resChan <- res
		}(c)
	}
	startFunc()
	go func() {
		wg.Wait()
		close(resChan)
	}()
	numErrs := 0
	for res := range resChan {
		cmdStr := res.cfgCmd.Exec
		if res.err != nil {
			color.Red("Error running '%s' (%s)", cmdStr, res.err)
			logger.Printf("Stdout:")
			color.Red(res.stdout.String())
			logger.Printf("Stderr:")
			color.Red(res.stderr.String())
			if res.crash {
				os.Exit(1)
			}
			numErrs++
		} else {
			logger.Printf("%s", cmdStr)
			color.Green("Success!")
		}
		logger.Printf("\n")
	}
	logger.Printf("\n\n")
	if numErrs > 0 {
		color.Red("found %d error(s)", numErrs)
		return fmt.Errorf("Errors found!")
	}
	color.Green("Everything worked!")
	return nil
}
