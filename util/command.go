package util

import (
	"bytes"
	"errors"
	"os/exec"
)

// RunCommand Run string content as a command with given input. Return command results as a string or an error.
func RunCommand(command string) (output string, err error) {
	var cmd *exec.Cmd
	var cmdOutput bytes.Buffer
	var cmdErrorOutput bytes.Buffer
	cmd.Stdout = &cmdOutput
	cmd.Stderr = &cmdErrorOutput

	cmd = exec.Command("/bin/sh", "-c", command)

	err = cmd.Run()
	if err == nil {
		err = errors.New(cmdErrorOutput.String())
		output = cmdOutput.String()
	}
	return
}
