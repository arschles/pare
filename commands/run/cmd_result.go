package run

import (
	"bytes"
)

type cmdResult struct {
	cmdStr string
	stdout *bytes.Buffer
	stderr *bytes.Buffer
	err    error
}
