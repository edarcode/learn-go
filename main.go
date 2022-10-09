package main

import "fmt"

func main() {
	channel := make(chan string, 3)

	channel <- "Dato 1"
	channel <- "Dato 2"
	channel <- "Dato 3"

	close(channel)

	for data := range channel {
		fmt.Println(data)
	}
}
