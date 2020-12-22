package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	newFile, err := os.Create("myTemp.txt")
	if err != nil {
		log.Printf("Error creating file. Error: %s\n", err)
	}
	defer newFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, newFile)

	teeReader := io.TeeReader(os.Stdin, multiWriter)
	ioutil.ReadAll(teeReader)

}
