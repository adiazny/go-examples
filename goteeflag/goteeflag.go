package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	multiWriter := io.MultiWriter(os.Stdout)
	writers := make([]io.Writer, 1)
	writers[0] = os.Stdout

	canAppend := flag.Bool("a", false, "Append the output to the files rather than overwriting them.")

	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {

		for i := 0; i < len(args); i++ {
			file := CreateFile(args[i], *canAppend)
			defer file.Close()
			writers = append(writers, file)
		}

		multiWriter = io.MultiWriter(writers...)
	}

	teeReader := io.TeeReader(os.Stdin, multiWriter)
	ioutil.ReadAll(teeReader)
}

// CreateFile returns a file with append or no append ability
func CreateFile(fileName string, canAppend bool) *os.File {
	fileMode := os.O_TRUNC

	if canAppend {
		fileMode = os.O_APPEND
	}

	file, err := os.OpenFile(fileName, fileMode|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	return file

}
