package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	envs := make(Environment)
	for _, file := range files {
		fullFileName, err := os.Open(fmt.Sprintf("%s/%s", dir, file.Name()))
		if err != nil {
			return nil, err
		}
		reader := bufio.NewReader(fullFileName)
		line, err := reader.ReadString('\n')
		if errors.Is(err, io.EOF) {
			err = nil
		}
		if err != nil {
			fullFileName.Close()
			return nil, err
		}
		if err = fullFileName.Close(); err != nil {
			return nil, err
		}

		line = string(bytes.ReplaceAll([]byte(line), []byte{0x00}, []byte("\n")))
		line = string(bytes.TrimRight([]byte(line), "\n\r\t "))

		var envVal EnvValue
		envVal.Value = line
		if len(line) == 0 {
			envVal.NeedRemove = true
		}
		envs[strings.ReplaceAll(file.Name(), "=", "")] = envVal
	}
	return envs, nil
}
