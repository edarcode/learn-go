package main

import "fmt"

func main() {
	var students [4]int

	for _, student := range students {
		fmt.Println(student)
	}
}
