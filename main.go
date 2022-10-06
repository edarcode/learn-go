package main

import (
	"fmt"
	"learn-go/utils"
)

func main() {
	addition, subtraction, multiplication, division := utils.GetBasicOperations(1, 0)
	fmt.Println(addition)
	fmt.Println(subtraction)
	fmt.Println(multiplication)
	fmt.Println(division)
}
