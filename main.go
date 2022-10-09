package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Dato 1"
	channel <- "Dato 2"

	fmt.Println("len:", len(channel))
	fmt.Println("cap:", cap(channel))
}
