package main

import "fmt"

func main() {
	channel := make(chan string, 3)

	channel <- "Dato 1"
	channel <- "Dato 2"

	fmt.Println(len(channel)) // 2
	fmt.Println(cap(channel)) // 3
}
