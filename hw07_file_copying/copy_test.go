package main

import (
	"errors"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("Offset bigger then file size", func(t *testing.T) {
		offset = 1024
		from = "./testdata/tesssst.txt"
		res := Copy(from, to, offset, limit)
		require.Equal(t, res, errors.New("offset exceeds file size"))
	})
	t.Run("Unsupported file", func(t *testing.T) {
		from = "/dev/urandom"
		res := Copy(from, to, offset, limit)
		require.Equal(t, res, errors.New("unsupported file"))
	})
	t.Run("Successful copying with limit and offset", func(t *testing.T) {
		from = "./testdata/t.csv"
		limit = 50
		offset = 10
		Copy(from, to, offset, limit)

		fileSrc, err := os.Open(from)
		if err != nil {
			return
		}
		defer fileSrc.Close()
		stSrc, err := fileSrc.Stat()
		if err != nil {
			return
		}
		sizeOldFile := stSrc.Size()

		fileDst, err := os.Open(to)
		if err != nil {
			return
		}
		defer fileDst.Close()
		stDst, err := fileDst.Stat()
		if err != nil {
			return
		}
		sizeNewFile := stDst.Size()
		require.Less(t, sizeNewFile, sizeOldFile)
	})
}
