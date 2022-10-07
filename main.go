package main

import (
	"fmt"
	"learn-go/structs"
)

func main() {
	edarcode := structs.Person{Name: "edarcode", Age: 26}
	lore := structs.Person{Name: "lorena", Age: 40}

	fmt.Println(edarcode.Greet()) // Hola soy edarcode y tengo 26 años
	fmt.Println(lore.Greet())     // Hola soy lorena y tengo 40 años
}
