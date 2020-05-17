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
		require.Equal(t, one, 0)
		two := RunCmd(nil, env)
		require.Equal(t, two, 0)
	})
	t.Run("CMD + null env", func(t *testing.T) {
		cmd := []string{"ipconfig"} //ifconfig
		r := RunCmd(cmd, nil)
		require.Equal(t, 1, r)
	})

}
