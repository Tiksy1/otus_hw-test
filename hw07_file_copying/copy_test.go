package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const fromPath = "testdata"

const toPath = "testdata/test"

const fromFile = "input.txt"

const toFile = "out.txt"

var fromPathTest = filepath.Join(fromPath, fromFile)

var toPathTest = filepath.Join(toPath, toFile)

func TestCopy(t *testing.T) {
	srcFile, err := os.Open(fromPathTest)
	require.NoError(t, err)
	err = os.Mkdir(toPath, 0o755)
	defer func() {
		err = srcFile.Close()
		os.RemoveAll(toPath)
		require.NoError(t, err)
	}()
	testCases := []struct {
		name   string
		offset int64
		limit  int64
	}{
		{"offset0_limit0", 0, 0},
		{"offset0_limit10", 0, 10},
		{"offset0_limit1000", 0, 1000},
		{"offset0_limit10000", 0, 10000},
		{"offset100_limit100", 100, 100},
		{"offset100_limit1000", 100, 1000},
		{"offset6000_limit1000", 6000, 1000},
		{"offset-1_limit-1", -1, -1},
		{"offset-1_limit0", -1, 0},
		{"offset0_limit-1", 0, -1},
		{"offset-1_limit1", -1, 1},
		{"offset1_limit-1", 1, -1},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err = Copy(srcFile.Name(), toPathTest, tc.offset, tc.limit)
			require.NoError(t, err)

			err = os.Remove(toPathTest)
			require.NoError(t, err)
		})
	}
}

func TestNoPermToFile(t *testing.T) {
	err := Copy("/dev/urandom", toFile, 0, 0)

	require.Equal(t, true, errors.Is(err, ErrUnsupportedFile))
}
