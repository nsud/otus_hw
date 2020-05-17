package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]string

var resList Environment

func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if resList == nil {
			resList = make(Environment)
		}
		fullPathFile := filepath.Join(dir, file.Name())
		if !file.IsDir() {
			if strings.Contains(file.Name(), "=") {
				//return nil, errors.New("Unexpected symbol")
				continue
			}
			//openFile, err := ioutil.ReadFile(fullPathFile)
			openFile, err := os.Open(fullPathFile)
			if err != nil {
				fmt.Println(errors.New("non-readable file"))
				continue
			}
			defer openFile.Close()
			st, err := openFile.Stat()
			if err != nil {
				return nil, err
				//fmt.Errorf("%v", err)
			}
			if st.Size() == 0 {
				//fmt.Errorf("Empty file %v", st.Name())
				emptF := errors.New("empty file")
				//delete(resList, st.Name())
				return nil, emptF
				//continue
			}
			openFile2, _ := ioutil.ReadFile(fullPathFile)
			text := string(bytes.ReplaceAll(openFile2, []byte("0x00"), []byte("\n")))
			readyToList := strings.TrimRight(text, " \t\n")

			resList[file.Name()] = readyToList
		}
	}
	return resList, err
}
