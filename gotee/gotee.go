package main

//TODO: add the -i option. Ignore the SIGINT signal.

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]
	var multiWriter io.Writer
	writers := make([]io.Writer, 1)
	writers[0] = os.Stdout
	argsLength := len(args)
	canAppend := false
	index := 0

	if argsLength > 0 {
		if args[0] == "-a" {
			canAppend = true
			index++
		}

		for i := index; i < argsLength; i++ {
			file := CreateFile(args[i], canAppend)
			defer file.Close()
			writers = append(writers, file)
		}

		multiWriter = io.MultiWriter(writers...)
	} else {
		multiWriter = io.MultiWriter(os.Stdout)
	}

	teeReader := io.TeeReader(os.Stdin, multiWriter)
	ioutil.ReadAll(teeReader)

}

// CreateFile returns a file with append or no append ability
func CreateFile(fileName string, canAppend bool) *os.File {
	var fileMode int

	if canAppend {
		fileMode = os.O_APPEND
	} else {
		fileMode = os.O_TRUNC
	}

	file, err := os.OpenFile(fileName, fileMode|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	return file

}
