package main

import "fmt"

func main() {
	var students = []string{"edarcode", "lore", "marian"}
	var teachers = []string{"jose", "pablo"}

	persons := append(students, teachers...)

	fmt.Println(persons)
}
