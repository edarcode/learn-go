package main

import (
	"fmt"
	"strings"
)

func main() {
	channel := make(chan string, 1)

	fmt.Println("HOLA") // HOLA
	go getUpperText("lore", channel)
	fmt.Println(<-channel) // LORE
}

func getUpperText(text string, channel chan<- string) {
	upperText := strings.ToUpper(text)
	channel <- upperText
}
