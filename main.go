package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value, err := strconv.Atoi("edar")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
}
