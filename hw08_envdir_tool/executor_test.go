package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("returns success exit code", func(t *testing.T) {
		execCmd := []string{"ls", "testdata/env"}
		exitCode := RunCmd(execCmd, nil)

		require.Equal(t, 0, exitCode)
	})

	t.Run("wrong directory", func(t *testing.T) {
		execCmd := []string{"ls", "testdata/env/wrongDir"}
		exitCode := RunCmd(execCmd, nil)

		require.Equal(t, -1, exitCode)
	})
}
