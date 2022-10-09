# Channel

## Select

Es posible que necesite leer todas las concurrencias en curso y decidir que hacer según el caso, podemos utilizar for y select.

Primero creamos las funciones que vamos a necesitar.

```
func getUsers(users chan<- []string) {
	users <- []string{"edarcode", "lore"}
}

func getSingers(singers chan<- []string) {
	singers <- []string{"shakira", "bad bunny"}
}
```

Creemos 2 canales y ejecutemos 2 **goroutine**.

```
func main() {
	users := make(chan []string, 1)
	singers := make(chan []string, 1)

	go getUsers(users)
	go getSingers(singers)
}
```

Procedemos a iterar la cantidad **goroutine** que tenemos. En los **case** debe especificar una varible e igualarla a la salida del canal o channel especifico.

```
case users := <-users:
  fmt.Println("usuarios:", users)
```

Iteramos y seleccionamos según el caso:

```
channelCount := 2
for i := 0; i < channelCount; i++ {
  select {
  case users := <-users:
    fmt.Println("usuarios:", users)
  case singers := <-singers:
    fmt.Println("cantantes:", singers)
  }
}
```

Vista general:

```
package main

import "fmt"

func main() {
	users := make(chan []string, 1)
	singers := make(chan []string, 1)

	go getUsers(users)
	go getSingers(singers)

	channelCount := 2
	for i := 0; i < channelCount; i++ {
		select {
		case users := <-users:
			fmt.Println("usuarios:", users)
		case singers := <-singers:
			fmt.Println("cantantes:", singers)
		}
	}
}

func getUsers(users chan<- []string) {
	users <- []string{"edarcode", "lore"}
}

func getSingers(singers chan<- []string) {
	singers <- []string{"shakira", "bad bunny"}
}
```
