package main

import "fmt"

func main() {
	fmt.Println("edarcode", 26)
	fmt.Printf("%s tiene %d años\n", "edarcode", 26)
	fmt.Printf("%T\n%T\n", "edarcode", 26)
	var msg string = fmt.Sprintf("%s tiene %d años\n", "edarcode", 26)
	fmt.Println(msg)

}
