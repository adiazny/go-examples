package main

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

	if argsLength > 0 {

		for i := 0; i < argsLength; i++ {
			file, err := os.Create(args[i])
			if err != nil {
				log.Printf("Error creating file. Error: %s\n", err)
			}
			defer file.Close()
			writers = append(writers, file)
		}

		multiWriter = io.MultiWriter(writers...)
	} else {
		multiWriter = io.MultiWriter(os.Stdout)
	}

	teeReader := io.TeeReader(os.Stdin, multiWriter)
	ioutil.ReadAll(teeReader)

	//TODO: add the -a option. Append the output to the files rather than overwriting them.
	//TODO: add the -i option. Ignore the SIGINT signal.
}
