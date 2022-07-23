package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("arguments count is incorrect")
	}
	env, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	exitCode := RunCmd(os.Args[2:], env)
	os.Exit(exitCode)
}
