package main

import "fmt"

func main() {
	var edarcode [4]int
	edarcode[0] = 1
	edarcode[1] = 2
	edarcode[2] = 3

	fmt.Println(edarcode)
	fmt.Println(len(edarcode))
	fmt.Println(cap(edarcode))
}
