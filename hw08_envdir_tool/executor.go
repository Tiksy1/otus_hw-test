package main

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) < 1 {
		return 1
	}
	for key, val := range env {
		if val.NeedRemove {
			_ = os.Unsetenv(key)
		} else {
			_ = os.Setenv(key, val.Value)
		}
	}

	execCmd := exec.Command(cmd[0], cmd[1:]...) // nolint:gosec
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		fmt.Println(err)
		return -1
	}
	return execCmd.ProcessState.ExitCode()
}
