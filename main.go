package main

import (
	"fmt"
)

func main() {
	things := []interface{}{"perro", 50, "gato", nil, 2.5}
	fmt.Println(things)    // [perro 50 gato <nil> 2.5]
	fmt.Println(things...) // perro 50 gato <nil> 2.5
}
