package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	num, err := strconv.Atoi("123hi45")
	if err != nil {
		log.Printf("Error converting: %+v", err)
	}

	fmt.Println(num)

}
