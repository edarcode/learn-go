# Channel

## Close

Permite cerrar un **channel**. En este code nos va a lanzar un err dado estamos intentando enviar un dato después de cerrar el **channel**.

```
package main

func main() {
	channel := make(chan string, 3)

	channel <- "Dato 1"
	channel <- "Dato 2"
	close(channel)
	channel <- "Dato 3"
}
```

Corrijamos el err e iteremos el channel con **range** para obtener sus datos

```
package main

import "fmt"

func main() {
	channel := make(chan string, 3)

	channel <- "Dato 1"
	channel <- "Dato 2"
	channel <- "Dato 3"

	close(channel)

	for data := range channel {
		fmt.Println(data)
	}
}
```
