package main

import (
	"log"
	"os"
)

var magicNumber = 3

func main() {
	args := os.Args
	if len(args) < magicNumber {
		log.Fatalf("You must use 3 arguments. For example, /path/to/env/dir command arg1 arg2")
	}
	pathEnvFile := args[1]
	commands := args[2:]

	envDir, err := ReadDir(pathEnvFile)
	if err != nil {
		log.Fatal(err)
	}

	resp := RunCmd(commands, envDir)
	if resp != 0 {
		log.Fatalf("Response code: %v", resp)
	}
}
