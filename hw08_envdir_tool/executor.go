package main

import (
	"os"
	"os/exec"
)

const (
	codeSuccess = 1
	codeFailure = 0
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return codeFailure
	}

	name, args := cmd[0], cmd[1:]
	proc := exec.Command(name, args...)
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	if env != nil {
		envs := make([]string, 0, len(env))
		for k, v := range env {
			val := k + "=" + v
			envs = append(envs, val)
		}

		proc.Env = envs
	}

	if err := proc.Run(); err != nil {
		//fmt.Println(err)
		return codeFailure
	}

	return codeSuccess
}
