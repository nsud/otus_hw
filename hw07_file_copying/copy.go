package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	stat, err := file.Stat()

	switch {
	case err != nil:
		return ErrUnsupportedFile
	case offset > stat.Size():
		return ErrOffsetExceedsFileSize
	case !stat.Mode().IsRegular():
		return ErrUnsupportedFile
	case offset != 0:
		_, err := file.Seek(offset, 0)
		if err != nil {
			return err
		}
		fallthrough
	case limit == 0:
		limit = stat.Size() - offset
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(file)

	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	for {
		c, err := io.CopyN(newFile, barReader, limit)
		if err == io.EOF || c == limit {
			break
		}
	}
	//newFile.Chmod(0644)
	bar.Finish()

	return nil
}
