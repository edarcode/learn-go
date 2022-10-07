package main

import (
	"fmt"
	"learn-go/structs"
)

func main() {
	edarcode := structs.CreateUser("edarcode", 26)
	fmt.Println(edarcode.GetName()) // edarcode
	fmt.Println(edarcode.GetAge())  //26
}
