package main

import (
	"bufio"
	"bytes"
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
				continue
			}
			openFile, err := os.Open(fullPathFile)
			if err != nil {
				return nil, err
			}
			defer openFile.Close()

			scanner := bufio.NewScanner(openFile)
			for scanner.Scan() {
				firstLine := scanner.Text()
				text := string(bytes.ReplaceAll([]byte(firstLine), []byte("0x00"), []byte("\n")))
				readyToList := strings.TrimRight(text, " \t\n")
				resList[file.Name()] = readyToList
				break
			}
			if err := scanner.Err(); err != nil {
				return nil, nil
			}
		}
	}
	return resList, err
}
