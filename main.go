package main

import "fmt"

func main() {
	edarcode := person{name: "edarcode", age: 26}

	fmt.Println(edarcode.name, edarcode.age)
}

type person struct {
	name string
	age  int
}
