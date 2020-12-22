package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	args := os.Args
	fmt.Printf("args type %T\n", args)

	for ix, val := range args {
		fmt.Printf("Argument %v: %s\n", ix, val)
	}

	newFile, err := os.Create("myTemp.txt")
	if err != nil {
		log.Printf("Error creating file. Error: %s\n", err)
	}
	defer newFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, newFile)

	teeReader := io.TeeReader(os.Stdin, multiWriter)
	ioutil.ReadAll(teeReader)

}
