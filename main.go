package main

import (
	"fmt"
	"learn-go/structs"
)

func main() {
	edarcode := structs.CreateUser("edarcode", 26, "client")
	fmt.Println(edarcode.GetName()) // edarcode
	fmt.Println(edarcode.GetAge())  // 26
	fmt.Println(edarcode.GetRole()) // client

	shakira := structs.CreateSinger("shakira isabel", 45, "shakira")
	fmt.Println(shakira.GetName())     // shakira isabel
	fmt.Println(shakira.GetAge())      // 45
	fmt.Println(shakira.GetNickname()) // shakira
}
