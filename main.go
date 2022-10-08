package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hola")

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("edarcode")
	}()

	wg.Wait()
}
