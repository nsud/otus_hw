package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadDir(t *testing.T) {
	t.Run("Incorrect path", func(t *testing.T) {
		_, err := ReadDir("./testdata7/")
		require.Error(t, err)
	})
	t.Run("Some file", func(t *testing.T) {
		_, err := ReadDir("./testdata/env/UNSET")
		require.Error(t, err)
	})
	t.Run("Empty file", func(t *testing.T) {
		f, err := ReadDir("./testdata/env/")
		if err != nil {
			return
		}
		require.Equal(t, "", f["UNSET"])

	})
	t.Run("Default case", func(t *testing.T) {
		f, err := ReadDir("./testdata/env/")
		if err != nil {
			return
		}
		require.Equal(t, "bar", f["BAR"])
		require.Equal(t, "", f["UNSET"])
	})

}
