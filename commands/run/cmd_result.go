package run

import (
	"bytes"

	"github.com/arschles/pare/config"
)

type cmdResult struct {
	cmdName string
	cfgCmd  *config.Command
	crash   bool
	stdout  *bytes.Buffer
	stderr  *bytes.Buffer
	err     error
}
