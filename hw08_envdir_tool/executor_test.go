package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("Empty cmd", func(t *testing.T) {
		cmd := []string{}
		env := make(Environment)
		env["ONE"] = "TWO"
		one := RunCmd(cmd, env)
		require.Equal(t, one, codeFailure)
		two := RunCmd(nil, env)
		require.Equal(t, two, codeFailure)
	})
	t.Run("CMD + null env", func(t *testing.T) {
		cmd := []string{"ifconfig"} //ifconfig
		r := RunCmd(cmd, nil)
		require.Equal(t, codeSuccess, r)
	})
	t.Run("default case", func(t *testing.T) {
		cmd := []string{}
		env := make(Environment)
		env["ONE"] = "TWO"
		one := RunCmd(cmd, env)
		require.Equal(t, one, codeFailure)
		two := RunCmd(nil, env)
		require.Equal(t, two, codeFailure)
	})

}
