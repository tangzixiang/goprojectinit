package utils

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestExecCommandWithMsg(t *testing.T) {

	stdoutMsg, stderrMsg, err := ExecCommandWithMsg("go", "mod")
	t.Logf("stdoutMsg: %v\n\n", stdoutMsg)
	t.Logf("stderrMsg: %v\n\n", stderrMsg)
	t.Logf("err : %v\n", err)

	if assert.IsType(t, &exec.ExitError{}, err) {
		t.Log("exec command result error type is *exec.ExitError")
	}
}

func TestExecCommand(t *testing.T) {
	t.Logf("go mode err : %v\n", ExecCommand("go", "mod"))
}
