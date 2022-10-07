package main

import (
	"fmt"
)

func main() {
	temperature := make(map[string]int)
	temperature["earth"] = 15
	temperature["mars"] = -65

	for key, value := range temperature {
		fmt.Println(key, value)
	}
	value, ok := temperature["xd"]
	fmt.Println(value, ok)
}
