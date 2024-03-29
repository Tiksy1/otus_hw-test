package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	envDir := "testdata/env"

	t.Run("valid data", func(t *testing.T) {
		if err := os.WriteFile("testdata/env/NEW=", []byte("new"), 0o755); err != nil {
			log.Fatal(err)
		}
		envs, err := ReadDir(envDir)

		require.NoError(t, err)
		require.Equal(t, "bar", envs["BAR"].Value)
		require.Equal(t, "", envs["EMPTY"].Value)
		require.Equal(t, "new", envs["NEW"].Value)
		if err := os.Remove("testdata/env/NEW="); err != nil {
			log.Fatal(err)
		}
	})
}
