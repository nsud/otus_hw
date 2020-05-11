package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	if err != nil {
		//os.Exit(1)
		return ErrUnsupportedFile
	}
	defer file.Close()

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

	fileN := toPath + stat.Name()
	newFile, err := os.Create(fileN)
	if err != nil {
		return err
	}
	defer newFile.Close()
	for {
		_, err := io.CopyN(newFile, barReader, limit)
		if err == io.EOF {
			break
		}
		return nil
	}
	newFile.Chmod(0644)
	bar.Finish()

	return nil
}
