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
	res := exec.Command(name, args...)
	res.Stdin = os.Stdin
	res.Stdout = os.Stdout
	res.Stderr = os.Stderr

	if env != nil {
		envs := make([]string, 0, len(env))
		for k, v := range env {
			if v == "" {
				os.Unsetenv(k)
				return codeFailure
			}
			val := k + "=" + v
			envs = append(envs, val)
		}
		res.Env = envs
	}

	if err := res.Run(); err != nil {
		return codeFailure
	}

	return codeSuccess
}
