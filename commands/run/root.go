package run

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/arschles/pare/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Root returns the command for the root of the 'pare run' command tree
func Root() *cobra.Command {
	return &cobra.Command{
		Use:     "run [target]",
		Aliases: []string{"r"},
		Short:   "run a target",
		Example: `pare run mytarget`,
		RunE:    run,
	}
}

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

	var wg sync.WaitGroup
	resChan := make(chan cmdResult)
	for i, cmdStr := range target.Commands {
		wg.Add(1)
		c := make(chan cmdResult)
		go func(cmdNum int, cmdStr string, resCh chan<- cmdResult) {
			defer close(resCh)
			if cmdStr == "" {
				return
			}
			spl := strings.Split(cmdStr, " ")
			first := spl[0]
			rest := spl[1:]
			cmd := exec.Command(first, rest...)
			stdoutBuf := new(bytes.Buffer)
			stderrBuf := new(bytes.Buffer)
			cmd.Stdout = stdoutBuf
			cmd.Stderr = stderrBuf
			err := cmd.Run()
			resCh <- cmdResult{
				cmdStr: cmdStr,
				stdout: stdoutBuf,
				stderr: stderrBuf,
				err:    err,
			}
		}(i, cmdStr, c)
		go func(ch <-chan cmdResult) {
			defer wg.Done()
			res := <-ch
			resChan <- res
		}(c)
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()
	numErrs := 0
	cmdNum := 0
	for res := range resChan {
		if res.err != nil {
			color.Red("Error running '%s' (%s)", res.cmdStr, res.err)
			logger.Printf("Stdout")
			color.Red(res.stdout.String())
			logger.Printf("Stderr")
			color.Red(res.stderr.String())
			numErrs++
		} else {
			logger.Printf("%d) %s", cmdNum, res.cmdStr)
			color.Green("Success!")
		}
		logger.Printf("\n")
		cmdNum++
	}
	logger.Printf("\n\n")
	if numErrs > 0 {
		color.Red("found %d error(s)", numErrs)
		return fmt.Errorf("Errors found!")
	}
	color.Green("Everything worked!")
	return nil
}
