package main

import (
	"fmt"
	"learn-go/structs"
)

func main() {
	edarcode := structs.Person{Name: "edarcode", Age: 26}
	fmt.Println(edarcode) // {edarcode 26}
	edarcode.SetName("edar")
	fmt.Println(edarcode) // {edarcode 26}
}
