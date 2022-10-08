# Channel

Permite pasar información entre **goroutines**, además, No es necesario indicar que espere la **goroutine** explícitamente. De esta forma el code está mucho mas limpio y logramos algo similar a lo que hicimos con los **wait groups**. Tener en cuenta que es un poco menos ineficiente que los **wg**.

Creemos un channel

```
channel := make(chan string, 1)
```

Luego una función que posteriormente ejecutaremos concurrentemente. Note que al final se debe especificar que va a recibir un **channel** indicando el tipo del mismo y el tipo de información que enviaran por el **channel**.

- **<-** Indica entrada del channel

```
func getUpperText(text string, channel chan<- string) {
  upperText := strings.ToUpper(text)
  channel <- upperText
}
```

Para consumir:

- **->** Indica salida del channel

```
func main() {
  channel := make(chan string, 1)

  fmt.Println("HOLA") // HOLA
  go getUpperText("lore", channel)
  fmt.Println(<-channel) // LORE
}
```

Vista general:

```
package main

import (
  "fmt"
  "strings"
)

func main() {
  channel := make(chan string, 1)

  fmt.Println("HOLA") // HOLA
  go getUpperText("lore", channel)
  fmt.Println(<-channel) // LORE
}

func getUpperText(text string, channel chan<- string) {
  upperText := strings.ToUpper(text)
  channel <- upperText
}
```
