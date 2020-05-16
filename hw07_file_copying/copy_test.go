package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("Offset bigger then file size", func(t *testing.T) {
		offset = 1024
		from = "./testdata/tesssst.txt"
		res := Copy(from, to, offset, limit)
		require.Equal(t, res, ErrOffsetExceedsFileSize)
	})
	t.Run("Unsupported file", func(t *testing.T) {
		from = "/dev/urandom"
		offset = 0
		res := Copy(from, to, offset, limit)
		require.Equal(t, res, ErrUnsupportedFile)
	})
	t.Run("Successful copying with limit and offset", func(t *testing.T) {
		from = "./testdata/t.csv"
		to = "/tmp/copy.csv"
		limit = 50
		offset = 10
		Copy(from, to, offset, limit)

		fileSrc, err := os.Open(from)
		if err != nil {
			return
		}
		defer fileSrc.Close()

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
		os.Remove(to)
		require.Equal(t, sizeNewFile, limit)

	})
	t.Run("Limit+offset > filesize", func(t *testing.T) {
		from = "./testdata/t.csv"
		to = "/tmp/"
		offset = 100
		limit = 100500
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
		os.Remove(to)
		require.Less(t, sizeNewFile, sizeOldFile-offset)
	})
}
