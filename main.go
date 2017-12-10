package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	// "github.com/spf13/cobra"
)

func main() {
	cfg := new(config)
	/*tomlMD*/ _, err := toml.DecodeFile("pare.toml", cfg)
	if err != nil {
		logger.Printf("Could not decode configuration file (%s)", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	resChan := make(chan cmdResult)
	for i, cmdStr := range cfg.Commands {
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
			logger.Printf("'%s' error (%s)", res.cmdStr, res.err)
			logger.Printf("stdout\n%s", res.stdout.String())
			logger.Printf("stderr\n%s", res.stderr.String())
			numErrs++
		} else {
			logger.Printf("%d) %s:\n\t%s", cmdNum, res.cmdStr, res.stdout.String())
		}
		cmdNum++
	}
	logger.Printf("found %d error(s)", numErrs)
	if numErrs > 0 {
		os.Exit(1)
	}
}
